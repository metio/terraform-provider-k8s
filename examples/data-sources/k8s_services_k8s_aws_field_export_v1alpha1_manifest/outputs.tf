output "manifests" {
  value = {
    "example" = data.k8s_services_k8s_aws_field_export_v1alpha1_manifest.example.yaml
  }
}
