data "k8s_endpoints_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
