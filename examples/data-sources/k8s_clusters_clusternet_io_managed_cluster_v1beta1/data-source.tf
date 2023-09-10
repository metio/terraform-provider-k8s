data "k8s_clusters_clusternet_io_managed_cluster_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
