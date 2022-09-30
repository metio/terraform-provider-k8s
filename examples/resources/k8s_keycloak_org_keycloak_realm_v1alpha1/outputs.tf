output "resources" {
  value = {
    "minimal" = k8s_keycloak_org_keycloak_realm_v1alpha1.minimal.yaml
  }
}
