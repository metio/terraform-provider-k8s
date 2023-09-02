data "k8s_opensearchservice_services_k8s_aws_domain_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
