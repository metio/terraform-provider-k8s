resource "k8s_gloo_solo_io_proxy_v1" "minimal" {
  metadata = {
    name = "test"
  }
}
