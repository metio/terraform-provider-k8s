resource "k8s_kuma_io_mesh_timeout_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
