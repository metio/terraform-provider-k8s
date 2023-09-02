output "manifests" {
  value = {
    "example" = data.k8s_ecr_services_k8s_aws_repository_v1alpha1_manifest.example.yaml
  }
}
