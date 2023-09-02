data "k8s_monitoring_coreos_com_thanos_ruler_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
