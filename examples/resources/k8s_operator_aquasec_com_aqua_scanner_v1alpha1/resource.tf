resource "k8s_operator_aquasec_com_aqua_scanner_v1alpha1" "minimal" {
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
    login = {
      host     = "some-host"
      username = "some-username"
      password = "some-password"
    }
  }
}
