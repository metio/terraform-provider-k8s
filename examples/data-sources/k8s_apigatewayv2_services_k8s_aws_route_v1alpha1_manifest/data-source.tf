data "k8s_apigatewayv2_services_k8s_aws_route_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
