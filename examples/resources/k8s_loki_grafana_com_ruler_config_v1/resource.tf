resource "k8s_loki_grafana_com_ruler_config_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}