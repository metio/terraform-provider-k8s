data "k8s_loki_grafana_com_loki_stack_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
