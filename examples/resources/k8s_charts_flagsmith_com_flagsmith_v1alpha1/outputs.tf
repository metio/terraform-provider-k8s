output "resources" {
  value = {
    "minimal" = k8s_charts_flagsmith_com_flagsmith_v1alpha1.minimal.yaml
    "example" = k8s_charts_flagsmith_com_flagsmith_v1alpha1.example.yaml
  }
}
