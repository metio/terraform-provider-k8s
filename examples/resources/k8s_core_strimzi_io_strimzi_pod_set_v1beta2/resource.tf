resource "k8s_core_strimzi_io_strimzi_pod_set_v1beta2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
