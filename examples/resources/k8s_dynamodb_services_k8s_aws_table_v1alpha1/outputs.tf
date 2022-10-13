output "resources" {
  value = {
    "minimal" = k8s_dynamodb_services_k8s_aws_table_v1alpha1.minimal.yaml
  }
}
