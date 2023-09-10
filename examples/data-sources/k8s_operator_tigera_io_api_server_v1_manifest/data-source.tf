data "k8s_operator_tigera_io_api_server_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
}
