output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_mesh_load_balancing_strategy_v1alpha1_manifest.example.yaml
  }
}
