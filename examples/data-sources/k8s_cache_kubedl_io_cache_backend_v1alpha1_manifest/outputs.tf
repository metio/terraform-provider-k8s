output "manifests" {
  value = {
    "example" = data.k8s_cache_kubedl_io_cache_backend_v1alpha1_manifest.example.yaml
  }
}
