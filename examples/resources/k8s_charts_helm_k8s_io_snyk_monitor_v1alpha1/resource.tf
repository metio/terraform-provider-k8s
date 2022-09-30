resource "k8s_charts_helm_k8s_io_snyk_monitor_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
