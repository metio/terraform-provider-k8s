data "k8s_longhorn_io_volume_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}