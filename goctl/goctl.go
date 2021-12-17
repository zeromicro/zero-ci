package goctl

import (
	"fmt"
	"time"

	"github.com/ory/dockertest/v3"
)

var successLabel = "Done"

func Run(res *dockertest.Resource) error {
	if err := runAPI(res); err != nil {
		return err
	}
	if err := runRPC(res); err != nil {
		return err
	}
	if err := runModel(res); err != nil {
		return err
	}
	if err := runKube(res); err != nil {
		return err
	}
	return runDocker(res)
}

func BuildAndRun(pool *dockertest.Pool, contextDir string) (*dockertest.Resource, error) {
	return pool.BuildAndRunWithBuildOptions(&dockertest.BuildOptions{
		ContextDir: contextDir,
	}, &dockertest.RunOptions{
		Name: "goctl",
		Tag:  "latest",
	})
}

func displayCmd(cmd []string) {
	fmt.Printf("\033[1;31;40m\nrunning cmd %v\033[0m\n", cmd)
	time.Sleep(time.Second)
}
