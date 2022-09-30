resource "k8s_wildfly_org_wild_fly_server_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    application_image = "some-image"
    replicas          = 7
  }
}
