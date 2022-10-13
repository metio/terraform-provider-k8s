resource "k8s_operator_aquasec_com_aqua_enforcer_v1alpha1" "minimal" {
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

resource "k8s_operator_aquasec_com_aqua_enforcer_v1alpha1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    common = {
      active_active     = true
      image_pull_secret = "aqua-registry"
    }
    deploy = {
      replicas = 3
      image = {
        registry    = "registry.aquasec.com"
        repository  = "enforcer"
        tag         = "<<IMAGE TAG>>"
        pull_policy = "IfNotPresent"
      }
    }
    gateway = {
      host = "some-host"
      port = 8080
    }
    infra = {
      requirements    = true
      service_account = "aqua-sa"
      version         = "2022.4"
    }
    token           = "some-token"
    run_as_non_root = false
  }
}
