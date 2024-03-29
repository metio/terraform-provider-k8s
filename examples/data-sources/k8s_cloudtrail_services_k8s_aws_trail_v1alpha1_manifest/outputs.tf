output "manifests" {
  value = {
    "example" = data.k8s_cloudtrail_services_k8s_aws_trail_v1alpha1_manifest.example.yaml
  }
}
