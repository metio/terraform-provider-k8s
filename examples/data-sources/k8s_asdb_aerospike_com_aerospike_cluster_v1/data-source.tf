data "k8s_asdb_aerospike_com_aerospike_cluster_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
