data "k8s_apisix_apache_org_apisix_plugin_config_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
