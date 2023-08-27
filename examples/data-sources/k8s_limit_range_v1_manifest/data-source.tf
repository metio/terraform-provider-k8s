data "k8s_limit_range_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    limits = [
      {
        type = "Pod"
        max = {
          cpu    = "200m"
          memory = "1024Mi"
        }
      },
      {
        type = "PersistentVolumeClaim"
        min = {
          storage = "24M"
        }
      },
      {
        type = "Container"
        default = {
          cpu    = "50m"
          memory = "24Mi"
        }
      },
    ]
  }
}
