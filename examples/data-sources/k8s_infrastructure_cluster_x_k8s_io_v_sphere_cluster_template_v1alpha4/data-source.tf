data "k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1alpha4" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
