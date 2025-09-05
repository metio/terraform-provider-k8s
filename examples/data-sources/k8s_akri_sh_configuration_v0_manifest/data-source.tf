data "k8s_akri_sh_configuration_v0_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
