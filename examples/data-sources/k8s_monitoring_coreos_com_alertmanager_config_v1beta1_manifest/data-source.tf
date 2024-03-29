data "k8s_monitoring_coreos_com_alertmanager_config_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {}
}
