output "manifests" {
  value = {
    "example" = data.k8s_apigatewayv2_services_k8s_aws_authorizer_v1alpha1_manifest.example.yaml
  }
}
