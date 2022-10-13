output "resources" {
  value = {
    "minimal" = k8s_lambda_services_k8s_aws_function_v1alpha1.minimal.yaml
  }
}
