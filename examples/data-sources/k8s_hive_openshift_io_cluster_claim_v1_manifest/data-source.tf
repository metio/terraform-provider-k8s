data "k8s_hive_openshift_io_cluster_claim_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_pool_name = "some-pool"
  }
}
