data "k8s_resources_teleport_dev_teleport_saml_connector_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
