output "manifests" {
  value = {
    "example" = data.k8s_route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest.example.yaml
  }
}
