output "manifests" {
  value = {
    "example" = data.k8s_objectbucket_io_object_bucket_claim_v1alpha1_manifest.example.yaml
  }
}
