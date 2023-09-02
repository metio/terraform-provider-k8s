resource "k8s_kuma_io_mesh_health_check_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
