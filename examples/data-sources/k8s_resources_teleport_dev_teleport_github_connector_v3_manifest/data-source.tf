data "k8s_resources_teleport_dev_teleport_github_connector_v3_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
