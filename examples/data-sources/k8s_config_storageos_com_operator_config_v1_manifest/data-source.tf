data "k8s_config_storageos_com_operator_config_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
