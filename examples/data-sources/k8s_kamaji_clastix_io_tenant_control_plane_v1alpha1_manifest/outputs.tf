output "manifests" {
  value = {
    "example" = data.k8s_kamaji_clastix_io_tenant_control_plane_v1alpha1_manifest.example.yaml
  }
}
