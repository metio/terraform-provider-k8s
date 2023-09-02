data "k8s_rules_kubeedge_io_rule_endpoint_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
