output "manifests" {
  value = {
    "example" = data.k8s_s3_snappcloud_io_s3_bucket_v1alpha1_manifest.example.yaml
  }
}
