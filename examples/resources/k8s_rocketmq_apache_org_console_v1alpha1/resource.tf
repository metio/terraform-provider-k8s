resource "k8s_rocketmq_apache_org_console_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    console_deployment = {}
  }
}
