data "k8s_apps_replica_set_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
    labels = {
      app  = "guestbook"
      tier = "frontend"
    }
  }
  spec = {
    replicas = 3
    selector = {
      match_labels = {
        tier = "frontend"
      }
    }
    template = {
      metadata = {
        labels = {
          tier = "frontend"
        }
      }
      spec = {
        containers = [
          {
            name  = "php-redis"
            image = "gcr.io/google_samples/gb-frontend:v3"
          }
        ]
      }
    }
  }
}
