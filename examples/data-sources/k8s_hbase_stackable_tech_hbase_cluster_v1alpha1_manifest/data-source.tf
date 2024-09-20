data "k8s_hbase_stackable_tech_hbase_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      hdfs_config_map_name      = "some-name"
      zookeeper_config_map_name = "some-name"
    }
    image = {}
  }
}
