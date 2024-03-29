output "manifests" {
  value = {
    "example" = data.k8s_registry_devfile_io_devfile_registry_v1alpha1_manifest.example.yaml
  }
}
