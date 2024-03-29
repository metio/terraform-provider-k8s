data "k8s_namespace_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
