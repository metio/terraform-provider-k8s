resource "k8s_scylla_scylladb_com_scylla_cluster_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
