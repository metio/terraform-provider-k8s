data "k8s_operator_authorino_kuadrant_io_authorino_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
