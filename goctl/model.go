package goctl

import "github.com/ory/dockertest/v3"

func runModel(res *dockertest.Resource) error {
	if err := genMysqlModel(res); err != nil {
		return err
	}
	if err := genMongoModel(res); err != nil {
		return err
	}
	return genPostgresqlModel(res)
}

func genMysqlModel(res *dockertest.Resource) error {
	return nil
}

func genPostgresqlModel(res *dockertest.Resource) error {
	return nil
}

func genMongoModel(res *dockertest.Resource) error {
	return nil
}
