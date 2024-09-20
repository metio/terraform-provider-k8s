data "k8s_sriovnetwork_openshift_io_sriov_network_node_policy_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
