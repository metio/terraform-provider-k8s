resource "k8s_agent_k8s_elastic_co_agent_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}

resource "k8s_agent_k8s_elastic_co_agent_v1alpha1" "example" {
  metadata = {
    name = "agent-sample"
  }
  spec = {
    version = "8.4.0"
    elasticsearch_refs = [
      {
        name = "elasticsearch-sample"
      }
    ]
    daemon_set = {}
    config = {
      inputs = [
        {
          name       = "system-1"
          revision   = 1
          type       = "system/metrics"
          use_output = "default"
          meta = {
            package = {
              name    = "system"
              version = "0.9.1"
            }
          }
          data_stream = {
            namespace = "default"
          }
          streams = [
            {
              id = "system/metrics-system.cpu"
              data_stream = {
                dataset = "system.cpu"
                type    = "metrics"
              }
              metricsets    = ["cpu"]
              "cpu.metrics" = ["percentages", "normalized_percentages"]
              period        = "10s"
            }
          ]
        }
      ]
    }
  }
}
