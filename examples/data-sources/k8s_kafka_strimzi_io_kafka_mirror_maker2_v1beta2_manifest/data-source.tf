data "k8s_kafka_strimzi_io_kafka_mirror_maker2_v1beta2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
