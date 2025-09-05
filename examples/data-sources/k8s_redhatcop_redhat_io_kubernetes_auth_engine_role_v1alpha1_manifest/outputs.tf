output "manifests" {
  value = {
    "example" = data.k8s_redhatcop_redhat_io_kubernetes_auth_engine_role_v1alpha1_manifest.example.yaml
  }
}
