resource "k8s_gloo_solo_io_settings_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
