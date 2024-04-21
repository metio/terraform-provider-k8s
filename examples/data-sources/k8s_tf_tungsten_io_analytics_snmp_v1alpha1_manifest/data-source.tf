data "k8s_tf_tungsten_io_analytics_snmp_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
