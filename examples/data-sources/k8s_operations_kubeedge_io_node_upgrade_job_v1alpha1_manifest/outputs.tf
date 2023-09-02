output "manifests" {
  value = {
    "example" = data.k8s_operations_kubeedge_io_node_upgrade_job_v1alpha1_manifest.example.yaml
  }
}
