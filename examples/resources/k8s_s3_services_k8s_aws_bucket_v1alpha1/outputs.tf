output "resources" {
  value = {
    "minimal" = k8s_s3_services_k8s_aws_bucket_v1alpha1.minimal.yaml
  }
}
