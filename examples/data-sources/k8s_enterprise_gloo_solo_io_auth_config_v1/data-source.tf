data "k8s_enterprise_gloo_solo_io_auth_config_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}