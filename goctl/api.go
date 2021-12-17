package goctl

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/ory/dockertest/v3"
)

var (
	showAPIDirCmd = []string{"tree", "./test-api"}
	genAPICmd     = []string{"goctl", "api", "go", "-api", "./test-api/test.api", "-dir", "./test-api", "-style", "gozero"}
	tidyAPICmd    = []string{"/bin/sh", "-c", "cd ./test-api && go mod tidy"}
	buildAPICmd   = []string{"/bin/sh", "-c", "cd ./test-api && go build -o test test.go"}
	runAPICmd     = []string{"/bin/sh", "-c", "cd ./test-api && ./test &"}
	curlAPICmd    = []string{"curl", "-i", "http://127.0.0.1:8888/test?name=zero"}
)

func runAPI(res *dockertest.Resource) error {
	if _, err := genAPI(res); err != nil {
		return err
	}
	_, err := runAPISrv(res)
	return err
}

func genAPI(res *dockertest.Resource) (string, error) {
	displayCmd(genAPICmd)
	b := make([]byte, 10240)
	buf := bytes.NewBuffer(b)
	if _, err := res.Exec(genAPICmd, dockertest.ExecOptions{StdOut: buf}); err != nil {
		return "", err
	}
	if !strings.Contains(buf.String(), successLabel) {
		return "", fmt.Errorf(buf.String())
	}
	displayCmd(showAPIDirCmd)
	if _, err := res.Exec(showAPIDirCmd, dockertest.ExecOptions{StdOut: buf}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func runAPISrv(res *dockertest.Resource) (string, error) {
	displayCmd(tidyAPICmd)
	b := make([]byte, 10240)
	buf := bytes.NewBuffer(b)
	if _, err := res.Exec(tidyAPICmd, dockertest.ExecOptions{StdOut: buf}); err != nil {
		return "", err
	}
	displayCmd(buildAPICmd)
	if _, err := res.Exec(buildAPICmd, dockertest.ExecOptions{StdOut: buf}); err != nil {
		return "", err
	}
	displayCmd(runAPICmd)
	if _, err := res.Exec(runAPICmd, dockertest.ExecOptions{StdOut: buf}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func curlSrv(res *dockertest.Resource) error {
	displayCmd(curlAPICmd)
	b := make([]byte, 10240)
	buf := bytes.NewBuffer(b)
	count := 0
	for {
		if count >= 12 {
			break
		}
		_, err := res.Exec(curlAPICmd, dockertest.ExecOptions{StdOut: buf})
		if err != nil {
			return err
		}
		if strings.Contains(buf.String(), "HTTP/1.1 200 OK") {
			return nil
		}
		time.Sleep(time.Second * 5)
		count++
	}
	return fmt.Errorf("api server start failure")
}
