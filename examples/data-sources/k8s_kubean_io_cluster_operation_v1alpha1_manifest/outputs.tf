output "manifests" {
  value = {
    "example" = data.k8s_kubean_io_cluster_operation_v1alpha1_manifest.example.yaml
  }
}
