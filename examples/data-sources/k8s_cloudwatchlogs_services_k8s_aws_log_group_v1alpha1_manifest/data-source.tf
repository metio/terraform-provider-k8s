data "k8s_cloudwatchlogs_services_k8s_aws_log_group_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
