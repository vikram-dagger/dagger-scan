package main

import (
	"context"
)

type Scan struct{}

func (m *Scan) Snyk(ctx context.Context, ctr *Container) (*Container, error) {
	return ctr, nil
}

func (ctr *Container) Snyk(ctx context.Context, snykToken *Secret) (*Container, error) {

	c := ctr.
		WithWorkdir("/tmp").
		WithExec([]string{"curl", "https://static.snyk.io/cli/latest/snyk-linux", "-o", "snyk"}).
		WithExec([]string{"chmod", "+x", "snyk"}).
		WithExec([]string{"mv", "./snyk", "/usr/local/bin"}).
		WithWorkdir("/src").
		WithSecretVariable("SNYK_TOKEN", snykToken).
		//WithExec([]string{"snyk", "monitor", "--all-projects", "--org=5e86b410-1a77-462a-a352-901a216fc3a6"})
		WithExec([]string{"snyk", "test"})

	return c, nil
}
