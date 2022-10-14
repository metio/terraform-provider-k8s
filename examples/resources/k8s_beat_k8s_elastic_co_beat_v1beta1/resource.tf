resource "k8s_beat_k8s_elastic_co_beat_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_beat_k8s_elastic_co_beat_v1beta1" "example" {
  metadata = {
    name = "heartbeat-sample"
  }
  spec = {
    type    = "heartbeat"
    version = "8.4.0"
    elasticsearch_ref = {
      name = "elasticsearch-sample"
    }
    config = {
      "heartbeat.monitors" = [
        {
          type     = "tcp"
          schedule = "@every 5s"
          hosts    = ["elasticsearch-sample-es-http.default.svc:9200"]
        }
      ]
    }
    deployment = {
      replicas = 1
      pod_template = {
        spec = {
          security_context = {
            run_as_user = 0
          }
          containers = []
        }
      }
    }
  }
}
