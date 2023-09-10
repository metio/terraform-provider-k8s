data "k8s_mq_services_k8s_aws_broker_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
