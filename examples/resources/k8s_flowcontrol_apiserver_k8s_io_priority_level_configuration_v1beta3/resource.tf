resource "k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    type = "Limited"
    limited = {
      assured_concurrency_shares = 125
    }
  }
}
