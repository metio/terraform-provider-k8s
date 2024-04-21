data "k8s_tf_tungsten_io_analytics_alarm_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
