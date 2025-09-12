data "k8s_secrets_stackable_tech_trust_store_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
