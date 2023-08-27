resource "k8s_pod_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
