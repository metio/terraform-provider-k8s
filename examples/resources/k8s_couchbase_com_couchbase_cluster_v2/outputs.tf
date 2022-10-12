output "resources" {
  value = {
    "minimal"   = k8s_couchbase_com_couchbase_cluster_v2.minimal.yaml
    "sha_image" = k8s_couchbase_com_couchbase_cluster_v2.sha_image.yaml
  }
}
