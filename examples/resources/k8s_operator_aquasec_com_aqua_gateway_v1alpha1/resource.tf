resource "k8s_operator_aquasec_com_aqua_gateway_v1alpha1" "minimal" {
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

resource "k8s_operator_aquasec_com_aqua_gateway_v1alpha1" "example" {
  metadata = {
    name      = "test"
    namespace = "test"
  }
  spec = {
    common = {
      active_active     = true
      image_pull_secret = "aqua-registry"
      split_db          = false
      database_secret = {
        name = "<<EXTERNAL DB PASSWORD SECRET NAME>>"
        key  = "<<EXTERNAL DB PASSWORD SECRET KEY>>"
      }
    }
    deploy = {
      replicas = 3
      service  = "ClusterIP"
      image = {
        registry    = "registry.aquasec.com"
        repository  = "enforcer"
        tag         = "<<IMAGE TAG>>"
        pull_policy = "IfNotPresent"
      }
    }
    infra = {
      requirements    = true
      service_account = "aqua-sa"
      version         = "2022.4"
    }
    external_db = {
      host     = "<<EXTERNAL DB IP OR DNS>>"
      port     = 12345
      username = "<<EXTERNAL DB USERNAME>>"
      password = "<<EXTERNAL DB PASSWORD (if secret does not exist)>>"
    }
    run_as_non_root = false
  }
}
