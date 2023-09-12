data "k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha3" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
