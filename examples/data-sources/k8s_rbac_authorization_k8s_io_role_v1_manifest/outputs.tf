output "manifests" {
  value = {
    "example" = data.k8s_rbac_authorization_k8s_io_role_v1_manifest.example.yaml
  }
}
