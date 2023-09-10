data "k8s_keycloak_org_keycloak_backup_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
