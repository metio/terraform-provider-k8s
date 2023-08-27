data "k8s_namespace_v1" "example" {
  metadata = {
    name = "some-name"

  }
}
