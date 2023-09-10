data "k8s_resources_teleport_dev_teleport_oidc_connector_v3_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
