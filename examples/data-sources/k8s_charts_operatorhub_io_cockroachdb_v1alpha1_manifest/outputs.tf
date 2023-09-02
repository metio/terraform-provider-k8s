output "manifests" {
  value = {
    "example" = data.k8s_charts_operatorhub_io_cockroachdb_v1alpha1_manifest.example.yaml
  }
}
