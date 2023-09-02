data "k8s_loki_grafana_com_recording_rule_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
