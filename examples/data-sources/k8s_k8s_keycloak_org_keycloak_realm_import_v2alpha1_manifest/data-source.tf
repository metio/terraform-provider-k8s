data "k8s_k8s_keycloak_org_keycloak_realm_import_v2alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
