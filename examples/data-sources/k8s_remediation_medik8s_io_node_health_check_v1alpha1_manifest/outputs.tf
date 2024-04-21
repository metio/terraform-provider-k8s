output "manifests" {
  value = {
    "example" = data.k8s_remediation_medik8s_io_node_health_check_v1alpha1_manifest.example.yaml
  }
}
