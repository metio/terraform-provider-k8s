output "manifests" {
  value = {
    "example" = data.k8s_rbacmanager_reactiveops_io_rbac_definition_v1beta1_manifest.example.yaml
  }
}
