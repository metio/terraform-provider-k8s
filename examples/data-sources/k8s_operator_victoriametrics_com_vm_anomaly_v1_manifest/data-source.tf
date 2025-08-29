data "k8s_operator_victoriametrics_com_vm_anomaly_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
