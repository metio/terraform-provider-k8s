data "k8s_namespace_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    annotations = {
      name = "example-annotation"
    }
    labels = {
      mylabel = "label-value"
    }
  }
}
