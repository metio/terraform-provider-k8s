resource "k8s_multicluster_x_k8s_io_work_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
