data "k8s_kafka_banzaicloud_io_kafka_user_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
