resource "k8s_source_toolkit_fluxcd_io_helm_repository_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
