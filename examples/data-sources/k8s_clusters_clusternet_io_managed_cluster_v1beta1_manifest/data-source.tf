data "k8s_clusters_clusternet_io_managed_cluster_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
