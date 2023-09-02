data "k8s_monitoring_coreos_com_service_monitor_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
