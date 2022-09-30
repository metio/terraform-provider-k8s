resource "k8s_flagger_app_alert_provider_v1beta1" "big" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
    labels = {
      "test" = "abc"
    }
    annotations = {
      "try" = "this"
    }
  }
  spec = {
    secret_ref = {
      name = "abc"
    }
    username = "user"
    type     = "some-type"
    proxy    = "localhost"
    channel  = "channel1"
    address  = "https://example.com"
  }
}

resource "k8s_flagger_app_alert_provider_v1beta1" "small" {
  metadata = {
    name = "test"
  }
  spec = {
    type    = "some-type"
    address = "https://example.com"
  }
}
