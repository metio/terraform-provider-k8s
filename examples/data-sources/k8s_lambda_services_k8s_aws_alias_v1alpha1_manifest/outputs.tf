output "manifests" {
  value = {
    "example" = data.k8s_lambda_services_k8s_aws_alias_v1alpha1_manifest.example.yaml
  }
}
