resource "k8s_replication_controller_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
