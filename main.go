package main

import (
	"flag"
	"log"

	"github.com/ory/dockertest/v3"
	"github.com/zeromicro/zero-ci/goctl"
)

var dir = flag.String("dir", "./", "the build context dir")

func main() {
	flag.Parse()

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal(err)
	}
	resource, err := goctlBuildAndRun(pool, *dir)
	if err != nil {
		log.Fatal(err)
	}
	if err = goctl.Run(resource); err != nil {
		log.Fatal(err)
	}
	if err = pool.Purge(resource); err != nil {
		log.Fatal(err)
	}
}

func goctlBuildAndRun(pool *dockertest.Pool, contextDir string) (*dockertest.Resource, error) {
	return pool.BuildAndRunWithBuildOptions(&dockertest.BuildOptions{
		ContextDir: contextDir,
	}, &dockertest.RunOptions{
		Name: "goctl",
		Tag:  "latest",
	})
}
