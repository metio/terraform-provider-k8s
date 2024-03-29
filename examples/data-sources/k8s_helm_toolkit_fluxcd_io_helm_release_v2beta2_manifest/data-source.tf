data "k8s_helm_toolkit_fluxcd_io_helm_release_v2beta2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
