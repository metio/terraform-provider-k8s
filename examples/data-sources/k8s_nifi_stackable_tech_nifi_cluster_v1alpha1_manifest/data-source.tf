data "k8s_nifi_stackable_tech_nifi_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      authentication = []
      sensitive_properties = {
        key_secret = "some-secret"
      }
      zookeeper_config_map_name = "some-name"
    }
    image = {}
  }
}
