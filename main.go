package main

import (
	"context"
)

type Scan struct{}

func (m *Scan) Snyk(ctx context.Context, ctr *Container) (*Container, error) {
	return ctr, nil
}

func (ctr *Container) Snyk(ctx context.Context, token string, path string) (*Container, error) {

	c := ctr.
		WithWorkdir("/tmp").
		WithExec([]string{"curl", "https://static.snyk.io/cli/latest/snyk-alpine", "-o", "snyk"}).
		WithExec([]string{"chmod", "+x", "snyk"}).
		WithExec([]string{"mv", "./snyk", "/usr/local/bin"}).
		WithWorkdir(path).
		WithEnvVariable("SNYK_TOKEN", token).
		WithExec([]string{"snyk", "test"})

	return c, nil
}
