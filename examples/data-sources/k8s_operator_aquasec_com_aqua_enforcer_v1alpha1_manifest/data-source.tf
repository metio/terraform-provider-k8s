data "k8s_operator_aquasec_com_aqua_enforcer_v1alpha1_manifest" "example" {
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
    gateway = {
      host = "some-host"
      port = 8080
    }
    infra = {
      requirements = true
    }
    token = "some-token"
  }
}
