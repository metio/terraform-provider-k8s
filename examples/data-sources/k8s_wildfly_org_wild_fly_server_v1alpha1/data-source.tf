data "k8s_wildfly_org_wild_fly_server_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
