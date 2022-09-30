resource "k8s_operator_aquasec_com_aqua_kube_enforcer_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    config = {}
  }
}
