data "k8s_service_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    selector = {
      app = "some-app"
    }
    session_affinity = "ClientIP"
    type             = "LoadBalancer"

    ports = [
      {
        port        = 8080
        target_port = 80
      },
    ]
  }
}
