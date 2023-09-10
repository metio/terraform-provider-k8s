data "k8s_rocketmq_apache_org_topic_transfer_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
