data "k8s_k8s_otterize_com_kafka_server_config_v1alpha2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
