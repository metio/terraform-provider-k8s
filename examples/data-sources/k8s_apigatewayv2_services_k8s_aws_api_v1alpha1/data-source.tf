data "k8s_apigatewayv2_services_k8s_aws_api_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
