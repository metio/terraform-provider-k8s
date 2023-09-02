data "k8s_source_toolkit_fluxcd_io_helm_repository_v1beta2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
