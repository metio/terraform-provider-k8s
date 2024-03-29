output "manifests" {
  value = {
    "example" = data.k8s_acmpca_services_k8s_aws_certificate_authority_activation_v1alpha1_manifest.example.yaml
  }
}
