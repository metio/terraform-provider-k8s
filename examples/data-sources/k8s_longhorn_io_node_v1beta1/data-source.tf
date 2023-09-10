data "k8s_longhorn_io_node_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
