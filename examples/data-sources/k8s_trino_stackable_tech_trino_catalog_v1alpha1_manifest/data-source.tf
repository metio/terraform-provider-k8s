data "k8s_trino_stackable_tech_trino_catalog_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
