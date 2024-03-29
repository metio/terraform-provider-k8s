output "manifests" {
  value = {
    "example" = data.k8s_cloudfront_services_k8s_aws_cache_policy_v1alpha1_manifest.example.yaml
  }
}
