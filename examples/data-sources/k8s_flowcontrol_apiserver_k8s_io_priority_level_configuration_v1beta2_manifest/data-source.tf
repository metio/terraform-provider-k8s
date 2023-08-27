data "k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    type = "Limited"
    limited = {
      assured_concurrency_shares = 125
    }
  }
}
