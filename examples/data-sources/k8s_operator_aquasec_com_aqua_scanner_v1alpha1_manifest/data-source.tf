data "k8s_operator_aquasec_com_aqua_scanner_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
