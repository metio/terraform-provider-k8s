data "k8s_apps_emqx_io_rebalance_v2beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
