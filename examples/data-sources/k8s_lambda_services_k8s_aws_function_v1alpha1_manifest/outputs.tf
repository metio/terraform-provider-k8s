output "manifests" {
  value = {
    "example" = data.k8s_lambda_services_k8s_aws_function_v1alpha1_manifest.example.yaml
  }
}