data "k8s_lambda_services_k8s_aws_event_source_mapping_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
