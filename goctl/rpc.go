package goctl

import (
	"os"

	"github.com/ory/dockertest/v3"
)

var (
	showRPCDirCmd = []string{"tree", "./test-rpc"}
	genRPCCmd     = []string{"goctl", "rpc", "proto", "-src", "./test-rpc/test.proto", "-dir", "./test-rpc"}
	runRPCCmd     = []string{"/bin/sh", "-c", "cd ./test-rpc && go mod tidy && nohup go run test.go &"}
)

func runRPC(res *dockertest.Resource) error {
	if err := genRPC(res); err != nil {
		return err
	}
	if err := runRPCSrv(res); err != nil {
		return err
	}
	return nil
}

func genRPC(res *dockertest.Resource) error {
	displayCmd(genRPCCmd)
	if _, err := res.Exec(genRPCCmd, dockertest.ExecOptions{
		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}); err != nil {
		return err
	}
	displayCmd(showRPCDirCmd)
	if _, err := res.Exec(showRPCDirCmd, dockertest.ExecOptions{
		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}); err != nil {
		return err
	}
	return nil
}

func runRPCSrv(res *dockertest.Resource) error {
	displayCmd(runRPCCmd)
	if _, err := res.Exec(runRPCCmd, dockertest.ExecOptions{
		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}); err != nil {
		return err
	}
	return nil
}
