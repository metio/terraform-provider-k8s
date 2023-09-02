output "manifests" {
  value = {
    "example" = data.k8s_scylla_scylladb_com_scylla_cluster_v1_manifest.example.yaml
  }
}
