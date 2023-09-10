data "k8s_kuma_io_mesh_rate_limit_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
