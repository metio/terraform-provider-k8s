data "k8s_source_toolkit_fluxcd_io_git_repository_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
