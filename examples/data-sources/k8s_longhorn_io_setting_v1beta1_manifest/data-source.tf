data "k8s_longhorn_io_setting_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  value = "some value"
}
