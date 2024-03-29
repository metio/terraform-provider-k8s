data "k8s_cloudtrail_services_k8s_aws_event_data_store_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
