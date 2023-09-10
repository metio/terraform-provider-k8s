data "k8s_kibana_k8s_elastic_co_kibana_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
