resource "k8s_config_map_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
