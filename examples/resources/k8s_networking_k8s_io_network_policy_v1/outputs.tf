output "resources" {
  value = {
    "minimal" = k8s_networking_k8s_io_network_policy_v1.minimal.yaml
    "example" = k8s_networking_k8s_io_network_policy_v1.example.yaml
  }
}
