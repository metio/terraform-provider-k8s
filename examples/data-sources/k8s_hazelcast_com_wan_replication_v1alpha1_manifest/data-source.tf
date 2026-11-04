data "k8s_hazelcast_com_wan_replication_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }

  spec = {
    endpoints           = "some-value"
    resources           = []
    target_cluster_name = "some-value"
  }
}
