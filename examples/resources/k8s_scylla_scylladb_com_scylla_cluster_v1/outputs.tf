output "resources" {
  value = {
    "minimal" = k8s_scylla_scylladb_com_scylla_cluster_v1.minimal.yaml
    "example" = k8s_scylla_scylladb_com_scylla_cluster_v1.example.yaml
  }
}
