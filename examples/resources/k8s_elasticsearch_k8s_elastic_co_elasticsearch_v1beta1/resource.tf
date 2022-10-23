resource "k8s_elasticsearch_k8s_elastic_co_elasticsearch_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_elasticsearch_k8s_elastic_co_elasticsearch_v1beta1" "example" {
  metadata = {
    name = "elasticsearch-sample"
  }
  spec = {
    version = "8.4.0"
    node_sets = [
      {
        name = "default"
        config = {
          "node.roles"            = "master"
          "node.attr.attr_name"   = "attr_value"
          "node.store.allow_mmap" = false
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
                name = "elasticsearch"
                resources = {
                  requests = {
                    memory = "4Gi"
                    cpu    = "1"
                  }
                  limits = {
                    memory = "4Gi"
                    cpu    = "2"
                  }
                }
              }
            ]
          }
        }
        count = 3
      }
    ]
  }
}
