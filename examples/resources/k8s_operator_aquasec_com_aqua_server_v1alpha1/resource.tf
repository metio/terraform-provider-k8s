resource "k8s_operator_aquasec_com_aqua_server_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    common = {
      active_active = true
    }
    deploy = {
      replicas = 3
    }
    infra = {
      requirements = true
    }
  }
}
