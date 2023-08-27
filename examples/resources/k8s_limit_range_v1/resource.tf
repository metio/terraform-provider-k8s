resource "k8s_limit_range_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
