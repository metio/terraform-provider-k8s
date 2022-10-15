resource "k8s_namespace_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_namespace_v1" "example" {
  metadata = {
    name = "terraform-example-namespace"
    annotations = {
      name = "example-annotation"
    }
    labels = {
      mylabel = "label-value"
    }
  }
}
