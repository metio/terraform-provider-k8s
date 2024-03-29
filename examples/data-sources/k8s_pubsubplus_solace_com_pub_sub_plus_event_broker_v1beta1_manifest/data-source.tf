data "k8s_pubsubplus_solace_com_pub_sub_plus_event_broker_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
