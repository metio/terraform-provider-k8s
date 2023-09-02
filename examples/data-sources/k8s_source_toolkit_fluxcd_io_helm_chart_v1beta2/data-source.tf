data "k8s_source_toolkit_fluxcd_io_helm_chart_v1beta2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
