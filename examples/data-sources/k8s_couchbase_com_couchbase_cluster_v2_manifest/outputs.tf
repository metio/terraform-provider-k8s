output "manifests" {
  value = {
    "example" = data.k8s_couchbase_com_couchbase_cluster_v2_manifest.example.yaml
  }
}
