data "k8s_operator_aquasec_com_aqua_kube_enforcer_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    config = {}
  }
}
