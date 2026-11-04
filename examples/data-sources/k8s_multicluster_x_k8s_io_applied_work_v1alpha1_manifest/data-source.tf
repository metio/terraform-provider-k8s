data "k8s_multicluster_x_k8s_io_applied_work_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }

  spec = {
    work_name      = "some-value"
    work_namespace = "some-value"
  }
}
