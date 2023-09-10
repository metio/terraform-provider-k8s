resource "k8s_cloudformation_linki_space_stack_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
