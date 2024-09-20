data "k8s_sriovnetwork_openshift_io_sriov_operator_config_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
