data "k8s_operator_cryostat_io_cryostat_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
