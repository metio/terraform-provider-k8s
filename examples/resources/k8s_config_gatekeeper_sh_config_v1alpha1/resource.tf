resource "k8s_config_gatekeeper_sh_config_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
