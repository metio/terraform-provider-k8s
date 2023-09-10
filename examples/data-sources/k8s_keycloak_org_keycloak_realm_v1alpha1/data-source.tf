data "k8s_keycloak_org_keycloak_realm_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
