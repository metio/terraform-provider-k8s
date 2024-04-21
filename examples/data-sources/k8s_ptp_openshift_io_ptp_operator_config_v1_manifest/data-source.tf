data "k8s_ptp_openshift_io_ptp_operator_config_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
