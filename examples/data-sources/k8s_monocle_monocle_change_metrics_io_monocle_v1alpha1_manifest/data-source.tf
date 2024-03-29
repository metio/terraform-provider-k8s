data "k8s_monocle_monocle_change_metrics_io_monocle_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
