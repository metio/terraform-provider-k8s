resource "k8s_rocketmq_apache_org_broker_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
