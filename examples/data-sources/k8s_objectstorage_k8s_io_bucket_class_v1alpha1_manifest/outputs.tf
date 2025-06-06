output "manifests" {
  value = {
    "example" = data.k8s_objectstorage_k8s_io_bucket_class_v1alpha1_manifest.example.yaml
  }
}
