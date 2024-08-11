terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "~> 2.-0"
    }
  }
}

provider "digitalocean" {
  token = var.do_token
}

resource "digitalocean_kubernetes_cluster" "analytics_cluster" {
  name    = "analytics-cluster"
  region  = "nyc1"
  version = "1.21.5-do.0"

  node_pool {
    name       = "worker-pool"
    size       = "s-2vcpu-2gb"
    node_count = 3
  }
}

resource "digitalocean_container_registry" "analytics_registry" {
  name                   = "analytics-registry"
  subscription_tier_slug = "starter"
}