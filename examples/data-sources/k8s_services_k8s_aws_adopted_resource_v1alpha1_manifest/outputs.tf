output "manifests" {
  value = {
    "example" = data.k8s_services_k8s_aws_adopted_resource_v1alpha1_manifest.example.yaml
  }
}
