output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_mesh_gateway_config_v1alpha1_manifest.example.yaml
  }
}
