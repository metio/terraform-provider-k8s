resource "k8s_persistent_volume_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
