resource "k8s_limit_range_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_limit_range_v1" "example" {
  metadata = {
    name = "test"
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
