data "k8s_services_k8s_aws_adopted_resource_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
