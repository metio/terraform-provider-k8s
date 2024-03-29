data "k8s_operator_tigera_io_tigera_status_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
