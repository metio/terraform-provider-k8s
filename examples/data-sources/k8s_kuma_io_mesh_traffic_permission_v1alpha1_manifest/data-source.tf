data "k8s_kuma_io_mesh_traffic_permission_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
