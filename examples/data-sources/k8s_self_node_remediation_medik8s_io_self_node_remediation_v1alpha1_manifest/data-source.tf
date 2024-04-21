data "k8s_self_node_remediation_medik8s_io_self_node_remediation_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
