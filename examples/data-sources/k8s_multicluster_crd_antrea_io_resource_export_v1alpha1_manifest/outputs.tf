output "manifests" {
  value = {
    "example" = data.k8s_multicluster_crd_antrea_io_resource_export_v1alpha1_manifest.example.yaml
  }
}
