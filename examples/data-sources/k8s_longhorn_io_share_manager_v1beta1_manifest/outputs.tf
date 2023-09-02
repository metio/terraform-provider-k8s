output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_share_manager_v1beta1_manifest.example.yaml
  }
}
