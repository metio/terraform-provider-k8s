output "resources" {
  value = {
    "minimal" = k8s_apps_replica_set_v1.minimal.yaml
    "example" = k8s_apps_replica_set_v1.example.yaml
  }
}
