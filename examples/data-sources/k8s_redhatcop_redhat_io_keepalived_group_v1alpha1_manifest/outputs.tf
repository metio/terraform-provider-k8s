output "manifests" {
  value = {
    "example" = data.k8s_redhatcop_redhat_io_keepalived_group_v1alpha1_manifest.example.yaml
  }
}
