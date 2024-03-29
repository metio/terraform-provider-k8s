output "manifests" {
  value = {
    "example" = data.k8s_secretsmanager_services_k8s_aws_secret_v1alpha1_manifest.example.yaml
  }
}
