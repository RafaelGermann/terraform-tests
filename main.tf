terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {
  host    = "npipe:////.//pipe//docker_engine"
}

resource "docker_image" "ubuntu" {
  name = "ubuntu:precise"
}

resource "docker_container" "ubuntu" {
  image = "ubuntu:latest"
  name  = "tutorial"
  command = ["/bin/bash", "-c", "tail -f /dev/null"]
}
