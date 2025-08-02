// A generated module for HelloWorld functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/hello-world/internal/dagger"
	"log"
)

type HelloWorld struct{}

// Returns a container that echoes whatever string argument is provided
func (m *HelloWorld) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *HelloWorld) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

func (m *HelloWorld) Build(ctx context.Context) *dagger.Container {
	return dag.Container().
		From("alpine:latest").
		WithEntrypoint([]string{"echo", "Hello, World!"})
}

func (m *HelloWorld) BuildAndPush(ctx context.Context, username string, password *dagger.Secret) error {
	registry := "ghcr.io"
	imageName := "lorenzofelletti/dagger-getting-started/hello-world"
	tags := []string{"latest", "v0.1.0"}

	container := m.Build(ctx).
		WithRegistryAuth(registry, username, password)

	for _, tag := range tags {
		addr, err := container.Publish(ctx, registry+"/"+imageName+":"+tag)
		if err != nil {
			return err
		}
		log.Printf("Published image: %s\n", addr)
	}

	return nil
}
