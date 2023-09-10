data "k8s_image_toolkit_fluxcd_io_image_policy_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
