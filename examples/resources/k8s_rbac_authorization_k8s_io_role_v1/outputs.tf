output "resources" {
  value = {
    "minimal" = k8s_rbac_authorization_k8s_io_role_v1.minimal.yaml
    "example" = k8s_rbac_authorization_k8s_io_role_v1.example.yaml
  }
}
