data "k8s_ipam_cluster_x_k8s_io_ip_address_claim_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
