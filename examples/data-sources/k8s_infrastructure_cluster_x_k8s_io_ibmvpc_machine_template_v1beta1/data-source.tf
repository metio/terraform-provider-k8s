data "k8s_infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
