data "k8s_image_toolkit_fluxcd_io_image_repository_v1beta2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
