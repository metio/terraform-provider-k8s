data "k8s_operator_victoriametrics_com_v_logs_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
