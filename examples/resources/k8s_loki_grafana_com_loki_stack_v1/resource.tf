resource "k8s_loki_grafana_com_loki_stack_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}