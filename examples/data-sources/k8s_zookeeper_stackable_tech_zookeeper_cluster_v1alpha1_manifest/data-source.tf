data "k8s_zookeeper_stackable_tech_zookeeper_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
