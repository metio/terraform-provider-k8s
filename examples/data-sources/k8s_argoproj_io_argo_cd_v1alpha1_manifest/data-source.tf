data "k8s_argoproj_io_argo_cd_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    controller = {
      resources = {
        limits = {
          cpu    = "2000m"
          memory = "2048Mi"
        }
        requests = {
          cpu    = "250m"
          memory = "1024Mi"
        }
      }
    }
    ha = {
      enabled = false
      resources = {
        limits = {
          cpu    = "500m"
          memory = "256Mi"
        }
        requests = {
          cpu    = "250m"
          memory = "128Mi"
        }
      }
    }
    redis = {
      resources = {
        limits = {
          cpu    = "500m"
          memory = "256Mi"
        }
        requests = {
          cpu    = "250m"
          memory = "128Mi"
        }
      }
    }
    repo = {
      resources = {
        limits = {
          cpu    = "1000m"
          memory = "512Mi"
        }
        requests = {
          cpu    = "250m"
          memory = "256Mi"
        }
      }
    }
    server = {
      resources = {
        limits = {
          cpu    = "500m"
          memory = "512Mi"
        }
        requests = {
          cpu    = "125m"
          memory = "128Mi"
        }
      }
      route = {
        enabled = true
      }
    }
    sso = {
      dex = {
        resources = {
          limits = {
            cpu    = "500m"
            memory = "256Mi"
          }
          requests = {
            cpu    = "250m"
            memory = "128Mi"
          }
        }
      }
      provider = "dex"
    }
  }
}
