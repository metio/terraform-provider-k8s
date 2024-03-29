output "manifests" {
  value = {
    "example" = data.k8s_redhatcop_redhat_io_patch_v1alpha1_manifest.example.yaml
  }
}
