data "k8s_wildfly_org_wild_fly_server_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    application_image = "some-image"
    replicas          = 7
  }
}
