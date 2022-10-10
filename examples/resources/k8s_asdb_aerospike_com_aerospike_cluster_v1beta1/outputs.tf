output "resources" {
  value = {
    "minimal" = k8s_asdb_aerospike_com_aerospike_cluster_v1beta1.minimal.yaml
    "example" = k8s_asdb_aerospike_com_aerospike_cluster_v1beta1.example.yaml
  }
}
