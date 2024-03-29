output "manifests" {
  value = {
    "example" = data.k8s_route53_services_k8s_aws_record_set_v1alpha1_manifest.example.yaml
  }
}
