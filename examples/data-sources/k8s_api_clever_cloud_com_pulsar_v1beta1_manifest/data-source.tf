data "k8s_api_clever_cloud_com_pulsar_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
