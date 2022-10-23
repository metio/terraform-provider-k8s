resource "k8s_networking_istio_io_destination_rule_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
