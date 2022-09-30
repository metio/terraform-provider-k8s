resource "k8s_operator_aquasec_com_aqua_csp_v1alpha1" "minimal" {
  metadata = {
    name = "test"
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
