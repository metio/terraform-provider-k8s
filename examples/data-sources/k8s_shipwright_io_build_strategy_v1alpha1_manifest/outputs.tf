output "manifests" {
  value = {
    "example" = data.k8s_shipwright_io_build_strategy_v1alpha1_manifest.example.yaml
  }
}
