data "k8s_app_lightbend_com_akka_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    replicas = 3
    selector = {
      match_labels = {
        "app.kubernetes.io/name" = "some-name"
      }
    }
    template = {
      spec = {
        containers = [
          {
            name  = "main"
            image = "lightbend-docker-registry.bintray.io/lightbend/akka-cluster-demo:1.1.0"
            readiness_probe = {
              http_get = {
                path = "/ready"
                port = "management"
              }
              period_seconds : 10
              failure_threshold : 10
              initial_delay_seconds : 20
            }
            liveness_probe = {
              http_get = {
                path = "/alive"
                port = "management"
              }
              period_seconds : 10
              failure_threshold : 10
              initial_delay_seconds : 20
            }
            ports = [
              {
                name           = "http"
                container_port = 8080
              },
              {
                name           = "remoting"
                container_port = 2552
              },
              {
                name           = "management"
                container_port = 8558
              },
            ]
          },
        ]
      }
    }
  }
}
