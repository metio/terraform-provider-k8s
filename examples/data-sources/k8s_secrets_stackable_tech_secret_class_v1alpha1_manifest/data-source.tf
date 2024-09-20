data "k8s_secrets_stackable_tech_secret_class_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    backend = {}
  }
}
