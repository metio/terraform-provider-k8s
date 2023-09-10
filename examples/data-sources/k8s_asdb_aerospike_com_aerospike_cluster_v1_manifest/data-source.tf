data "k8s_asdb_aerospike_com_aerospike_cluster_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    aerospike_config = {}
    image            = "aerospike/aerospike-server-enterprise:6.1.0.1"
    size             = 123
  }
}
