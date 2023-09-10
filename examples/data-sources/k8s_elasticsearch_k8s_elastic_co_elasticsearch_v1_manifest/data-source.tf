data "k8s_elasticsearch_k8s_elastic_co_elasticsearch_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    version = "8.4.0"
    node_sets = [
      {
        name = "default"
        config = {
          "node.roles"            = ["master", "data"]
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
