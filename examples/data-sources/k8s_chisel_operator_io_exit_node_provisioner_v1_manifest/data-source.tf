data "k8s_chisel_operator_io_exit_node_provisioner_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
