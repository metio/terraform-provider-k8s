data "k8s_trino_stackable_tech_trino_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      catalog_label_selector = {}
    }
    image = {}
  }
}
