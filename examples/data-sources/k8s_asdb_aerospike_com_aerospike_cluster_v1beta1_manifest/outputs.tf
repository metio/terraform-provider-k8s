output "manifests" {
  value = {
    "example" = data.k8s_asdb_aerospike_com_aerospike_cluster_v1beta1_manifest.example.yaml
  }
}
