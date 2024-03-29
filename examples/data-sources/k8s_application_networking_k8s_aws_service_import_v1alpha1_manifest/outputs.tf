output "manifests" {
  value = {
    "example" = data.k8s_application_networking_k8s_aws_service_import_v1alpha1_manifest.example.yaml
  }
}
