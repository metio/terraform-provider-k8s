resource "k8s_kuma_io_mesh_tcp_route_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
