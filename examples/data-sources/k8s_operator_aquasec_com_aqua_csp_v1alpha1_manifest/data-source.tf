data "k8s_operator_aquasec_com_aqua_csp_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    gateway = {
      replicas = 3
    }
    server = {
      replicas = 7
    }
  }
}
