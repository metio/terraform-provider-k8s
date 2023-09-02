output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_mesh_proxy_patch_v1alpha1_manifest.example.yaml
  }
}
