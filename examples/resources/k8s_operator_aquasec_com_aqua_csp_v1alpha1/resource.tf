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

resource "k8s_operator_aquasec_com_aqua_csp_v1alpha1" "example" {
  metadata = {
    name      = "aqua"
    namespace = "aqua"
  }
  spec = {
    infra = {
      service_account = "aqua-sa"
      namespace       = "aqua"
      version         = "2022.4"
      requirements    = true
    }
    common = {
      active_active     = true
      image_pull_secret = "aqua-registry"
      db_disk_size      = 10
      database_secret = {
        key  = "db-password"
        name = "aqua-database-password"
      }
    }
    database = {
      replicas = 1
      service  = "ClusterIP"
      image = {
        registry    = "registry.aquasec.com"
        repository  = "database"
        tag         = "<<IMAGE TAG>>"
        pull_policy = "Always"
      }
    }
    gateway = {
      replicas = 1
      service  = "ClusterIP"
      image = {
        registry    = "registry.aquasec.com"
        repository  = "gateway"
        tag         = "<<IMAGE TAG>>"
        pull_policy = "Always"
      }
    }
    server = {
      replicas = 1
      service  = "LoadBalancer"
      image = {
        registry    = "registry.aquasec.com"
        repository  = "gateway"
        tag         = "<<IMAGE TAG>>"
        pull_policy = "Always"
      }
    }
    route           = true
    run_as_non_root = false
  }
}
