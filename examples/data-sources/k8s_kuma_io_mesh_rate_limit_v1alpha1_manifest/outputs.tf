output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_mesh_rate_limit_v1alpha1_manifest.example.yaml
  }
}
