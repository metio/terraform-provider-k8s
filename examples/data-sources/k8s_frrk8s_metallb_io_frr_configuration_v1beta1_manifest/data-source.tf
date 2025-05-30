data "k8s_frrk8s_metallb_io_frr_configuration_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
