data "k8s_kubean_io_cluster_operation_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }

  spec = {
    action      = "some-value"
    action_type = "some-value"
    cluster     = "some-value"
    image       = "some-value"
  }
}
