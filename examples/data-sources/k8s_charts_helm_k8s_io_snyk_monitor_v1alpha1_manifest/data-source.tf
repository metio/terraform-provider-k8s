data "k8s_charts_helm_k8s_io_snyk_monitor_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
