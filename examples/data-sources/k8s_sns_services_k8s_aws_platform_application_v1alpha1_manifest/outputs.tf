output "manifests" {
  value = {
    "example" = data.k8s_sns_services_k8s_aws_platform_application_v1alpha1_manifest.example.yaml
  }
}
