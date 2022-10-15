output "resources" {
  value = {
    "minimal" = k8s_replication_controller_v1.minimal.yaml
    "example" = k8s_replication_controller_v1.example.yaml
  }
}
