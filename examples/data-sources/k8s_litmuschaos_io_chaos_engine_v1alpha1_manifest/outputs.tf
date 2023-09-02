output "manifests" {
  value = {
    "example" = data.k8s_litmuschaos_io_chaos_engine_v1alpha1_manifest.example.yaml
  }
}
