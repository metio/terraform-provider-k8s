data "k8s_cloudwatch_services_k8s_aws_metric_alarm_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
