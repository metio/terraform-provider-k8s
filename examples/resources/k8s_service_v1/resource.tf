resource "k8s_service_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_service_v1" "example" {
  metadata = {
    name = "test"
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
