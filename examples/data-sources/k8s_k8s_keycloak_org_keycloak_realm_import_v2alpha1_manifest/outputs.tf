output "manifests" {
  value = {
    "example" = data.k8s_k8s_keycloak_org_keycloak_realm_import_v2alpha1_manifest.example.yaml
  }
}
