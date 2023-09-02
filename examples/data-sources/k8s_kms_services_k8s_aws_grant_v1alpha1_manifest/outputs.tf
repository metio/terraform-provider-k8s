output "manifests" {
  value = {
    "example" = data.k8s_kms_services_k8s_aws_grant_v1alpha1_manifest.example.yaml
  }
}
