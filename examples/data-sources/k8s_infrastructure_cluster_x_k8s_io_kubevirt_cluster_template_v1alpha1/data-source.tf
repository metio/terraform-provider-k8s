data "k8s_infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
