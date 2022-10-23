resource "k8s_kibana_k8s_elastic_co_kibana_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_kibana_k8s_elastic_co_kibana_v1beta1" "example" {
  metadata = {
    name = "kibana-sample"
  }
  spec = {
    version = "8.4.0"
    count   = 1
    elasticsearch_ref = {
      name = "elasticsearch-sample"
    }
    pod_template = {
      metadata = {
        labels = {
          foo = "bar"
        }
      }
      spec = {
        containers = [
          {
            name = "kibana"
            resources = {
              requests = {
                memory = "1Gi"
                cpu    = "0.5"
              }
              limits = {
                memory = "2Gi"
                cpu    = "2"
              }
            }
          }
        ]
      }
    }
  }
}
