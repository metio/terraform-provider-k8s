data "k8s_prometheusservice_services_k8s_aws_alert_manager_definition_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
