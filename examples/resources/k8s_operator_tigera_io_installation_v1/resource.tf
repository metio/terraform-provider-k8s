resource "k8s_operator_tigera_io_installation_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {

  }
}
