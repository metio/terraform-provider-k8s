resource "k8s_keycloak_org_keycloak_client_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
