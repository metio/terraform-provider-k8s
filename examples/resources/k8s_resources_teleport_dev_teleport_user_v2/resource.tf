resource "k8s_resources_teleport_dev_teleport_user_v2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
