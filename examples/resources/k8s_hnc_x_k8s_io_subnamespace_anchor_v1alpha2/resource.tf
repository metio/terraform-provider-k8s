resource "k8s_hnc_x_k8s_io_subnamespace_anchor_v1alpha2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
