output "manifests" {
  value = {
    "example" = data.k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3_manifest.example.yaml
  }
}
