data "k8s_operator_cryostat_io_cryostat_v1beta2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
