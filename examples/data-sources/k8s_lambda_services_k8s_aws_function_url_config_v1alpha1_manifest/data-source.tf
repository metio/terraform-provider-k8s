data "k8s_lambda_services_k8s_aws_function_url_config_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
