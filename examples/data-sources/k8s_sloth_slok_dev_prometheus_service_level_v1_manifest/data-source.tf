data "k8s_sloth_slok_dev_prometheus_service_level_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
