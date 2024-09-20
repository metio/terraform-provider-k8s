data "k8s_sriovnetwork_openshift_io_sriov_ib_network_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
