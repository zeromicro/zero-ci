package goctl

import (
	"os"

	"github.com/ory/dockertest/v3"
)

var (
	showAPIDirCmd = []string{"tree", "./test-api"}
	genAPICmd     = []string{"goctl", "api", "go", "-api", "./test-api/test.api", "-dir", "./test-api", "-style", "gozero"}
	runAPICmd     = []string{"/bin/sh", "-c", "cd ./test-api && go mod tidy && nohup go run test.go &"}
)

func runAPI(res *dockertest.Resource) error {
	if err := genAPI(res); err != nil {
		return err
	}
	if err := runAPISrv(res); err != nil {
		return err
	}
	return nil
}

func genAPI(res *dockertest.Resource) error {
	displayCmd(genAPICmd)
	if _, err := res.Exec(genAPICmd, dockertest.ExecOptions{
		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}); err != nil {
		return err
	}
	displayCmd(showAPIDirCmd)
	if _, err := res.Exec(showAPIDirCmd, dockertest.ExecOptions{
		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}); err != nil {
		return err
	}
	return nil
}

func runAPISrv(res *dockertest.Resource) error {
	displayCmd(runAPICmd)
	if _, err := res.Exec(runAPICmd, dockertest.ExecOptions{
		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}); err != nil {
		return err
	}
	return nil
}
