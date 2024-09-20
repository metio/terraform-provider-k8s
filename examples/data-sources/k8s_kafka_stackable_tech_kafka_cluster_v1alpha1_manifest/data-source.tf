data "k8s_kafka_stackable_tech_kafka_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      zookeeper_config_map_name = "some-name"
    }
    image = {}
  }
}
