data "k8s_apicodegen_apimatic_io_api_matic_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
