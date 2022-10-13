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

resource "k8s_operator_aquasec_com_aqua_scanner_v1alpha1" "example" {
  metadata = {
    name      = "test"
    namespace = "test"
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
        repository  = "scanner"
        tag         = "<<IMAGE TAG>>"
        pull_policy = "IfNotPresent"
      }
    }
    infra = {
      requirements    = true
      service_account = "aqua-sa"
      version         = "2022.4"
    }
    login = {
      host     = "some-host"
      username = "some-username"
      password = "some-password"
      token    = ""
    }
    run_as_non_root = false
  }
}
