data "k8s_apps_emqx_io_emqx_plugin_v1beta4_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
