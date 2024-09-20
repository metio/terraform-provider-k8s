data "k8s_superset_stackable_tech_superset_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      credentials_secret = "some-secret"
    }
    image = {}
  }
}
