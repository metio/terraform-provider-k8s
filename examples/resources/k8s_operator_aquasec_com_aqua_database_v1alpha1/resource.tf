resource "k8s_operator_aquasec_com_aqua_database_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    common = {
      active_active = true
    }
    deploy = {
      replicas = 12
    }
    disk_size = 120
    infra = {
      requirements = true
    }
  }
}

resource "k8s_operator_aquasec_com_aqua_database_v1alpha1" "example" {
  metadata = {
    name      = "test"
    namespace = "test"
  }
  spec = {
    common = {
      active_active     = true
      image_pull_secret = "aqua-registry"
      split_db          = false
    }
    deploy = {
      replicas = 12
      service  = "ClusterIP"
      image = {
        registry    = "registry.aquasec.com"
        repository  = "database"
        tag         = "<<IMAGE TAG>>"
        pull_policy = "Always"
      }
    }
    disk_size       = 120
    run_as_non_root = false
    infra = {
      requirements    = true
      service_account = "aqua-sa"
      version         = "2022.4"
    }
  }
}
