data "k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_v1alpha4_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
