data "k8s_charts_flagsmith_com_flagsmith_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    api = {
      db_waiter = {
        image = {
          image_pull_policy = "IfNotPresent"
          repository        = "willwill/wait-for-it"
          tag               = "latest"
        }
      }
      image = {
        image_pull_policy = "IfNotPresent"
        repository        = "flagsmith/flagsmith-api"
        tag               = "2.8"
      }
      replicacount = 1
      resources = {
        limits = {
          cpu    = "300m"
          memory = "300Mi"
        }
        requests = {
          cpu    = "300m"
          memory = "300Mi"
        }
      }
    }
    database_external = {
      database = "dummy_db_name"
      enabled  = true
      host     = "dummy_db_host"
      password = "dummy_db_password"
      port     = 5432
      type     = "postgres"
      url      = "postgres://dummy_db_user:dummy_db_password@dummy_db_host:5432/dummy_db_name"
      username = "dummy_db_user"
    }
    frontend = {
      image = {
        image_pull_policy = "IfNotPresent"
        repository        = "flagsmith/flagsmith-frontend"
        tag               = "2.8"
      }
      replicacount = 1
      resources = {
        limits = {
          cpu    = "500m"
          memory = "500Mi"
        }
        requests = {
          cpu    = "300m"
          memory = "300Mi"
        }
      }
    }
  }
}
