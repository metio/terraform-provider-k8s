output "manifests" {
  value = {
    "example" = data.k8s_ecr_services_k8s_aws_pull_through_cache_rule_v1alpha1_manifest.example.yaml
  }
}
