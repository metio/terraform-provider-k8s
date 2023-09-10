data "k8s_gloo_solo_io_upstream_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
