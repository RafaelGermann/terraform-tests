package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
)

func TestDockerUbuntu(t *testing.T) {
	tag := "ubuntu:latest"
	containerName := "terratest-" + random.UniqueId()

	out := docker.Run(t, tag, &docker.RunOptions{
		Name:    containerName,
		Command: []string{"echo", "Hello, World"},
	})

	// Assert that the output contains the expected result
	assert.Contains(t, out, "Hello, World")
}
