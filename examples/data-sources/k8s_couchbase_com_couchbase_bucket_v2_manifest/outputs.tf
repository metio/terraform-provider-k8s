output "manifests" {
  value = {
    "example" = data.k8s_couchbase_com_couchbase_bucket_v2_manifest.example.yaml
  }
}
