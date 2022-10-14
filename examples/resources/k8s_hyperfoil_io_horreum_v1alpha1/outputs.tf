output "resources" {
  value = {
    "minimal" = k8s_hyperfoil_io_horreum_v1alpha1.minimal.yaml
    "example" = k8s_hyperfoil_io_horreum_v1alpha1.example.yaml
  }
}
