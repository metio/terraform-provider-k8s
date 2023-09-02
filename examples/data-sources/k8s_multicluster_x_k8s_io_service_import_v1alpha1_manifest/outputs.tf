output "manifests" {
  value = {
    "example" = data.k8s_multicluster_x_k8s_io_service_import_v1alpha1_manifest.example.yaml
  }
}
