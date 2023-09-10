data "k8s_flagger_app_alert_provider_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    secret_ref = {
      name = "abc"
    }
    username = "user"
    type     = "rocket"
    proxy    = "localhost"
    channel  = "channel1"
    address  = "https://example.com"
  }
}
