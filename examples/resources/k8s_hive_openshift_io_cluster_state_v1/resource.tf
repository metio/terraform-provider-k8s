resource "k8s_hive_openshift_io_cluster_state_v1" "minimal" {
  metadata = {
    name = "test"
  }
}
