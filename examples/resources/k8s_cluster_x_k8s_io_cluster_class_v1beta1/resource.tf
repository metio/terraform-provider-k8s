resource "k8s_cluster_x_k8s_io_cluster_class_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
