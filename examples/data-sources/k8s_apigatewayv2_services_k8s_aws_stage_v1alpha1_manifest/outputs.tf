output "manifests" {
  value = {
    "example" = data.k8s_apigatewayv2_services_k8s_aws_stage_v1alpha1_manifest.example.yaml
  }
}
