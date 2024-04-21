data "k8s_ptp_openshift_io_node_ptp_device_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
