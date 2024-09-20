data "k8s_authentication_stackable_tech_authentication_class_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    protocol = {}
  }
}
