output "cluster_endpoint" {
  value = digitalocean_kubernetes_cluster.analytics_cluster.endpoint
}

output "registry_endpoint" {
  value = digitalocean_container_registry.analytics_registry.server_url
}