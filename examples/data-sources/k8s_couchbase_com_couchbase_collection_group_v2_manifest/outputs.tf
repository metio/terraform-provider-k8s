output "manifests" {
  value = {
    "example" = data.k8s_couchbase_com_couchbase_collection_group_v2_manifest.example.yaml
  }
}
