output "resources" {
  value = {
    "minimal" = k8s_charts_operatorhub_io_cockroachdb_v1alpha1.minimal.yaml
    "example" = k8s_charts_operatorhub_io_cockroachdb_v1alpha1.example.yaml
  }
}
