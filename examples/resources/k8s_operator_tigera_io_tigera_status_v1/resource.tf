resource "k8s_operator_tigera_io_tigera_status_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {

  }
}
