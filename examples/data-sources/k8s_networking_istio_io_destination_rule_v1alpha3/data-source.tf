data "k8s_networking_istio_io_destination_rule_v1alpha3" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
