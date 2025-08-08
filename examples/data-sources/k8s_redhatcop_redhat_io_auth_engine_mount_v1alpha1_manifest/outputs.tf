output "manifests" {
  value = {
    "example" = data.k8s_redhatcop_redhat_io_auth_engine_mount_v1alpha1_manifest.example.yaml
  }
}
