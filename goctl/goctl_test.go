package goctl

import (
	"testing"

	"github.com/ory/dockertest/v3"
)

var (
	resource *dockertest.Resource
)

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}
	if resource, err = BuildAndRun(pool, ".."); err != nil {
		panic(err)
	}
	m.Run()
	_ = pool.Purge(resource)
}
