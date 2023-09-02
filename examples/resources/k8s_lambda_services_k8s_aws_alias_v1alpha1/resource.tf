resource "k8s_lambda_services_k8s_aws_alias_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
