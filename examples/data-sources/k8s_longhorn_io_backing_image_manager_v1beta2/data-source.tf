data "k8s_longhorn_io_backing_image_manager_v1beta2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}