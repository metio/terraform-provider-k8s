output "manifests" {
  value = {
    "example" = data.k8s_elasticache_services_k8s_aws_user_v1alpha1_manifest.example.yaml
  }
}
