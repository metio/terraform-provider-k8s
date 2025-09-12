data "k8s_operator_victoriametrics_com_vl_single_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
