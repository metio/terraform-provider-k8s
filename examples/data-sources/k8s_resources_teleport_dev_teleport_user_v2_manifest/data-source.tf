data "k8s_resources_teleport_dev_teleport_user_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
