resource "k8s_hive_openshift_io_cluster_claim_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    cluster_pool_name = "some-pool"
  }
}
