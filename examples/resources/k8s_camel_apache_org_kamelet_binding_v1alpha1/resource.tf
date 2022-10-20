resource "k8s_camel_apache_org_kamelet_binding_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_camel_apache_org_kamelet_binding_v1alpha1" "example" {
  metadata = {
    name = "telegram-text-source-to-channel"
  }
  spec = {
    source = {
      ref = {
        kind        = "Kamelet"
        api_version = "camel.apache.org/v1alpha1"
        name        = "telegram-text-source"
      }
      properties = {
        botToken = "the-token-here"
      }
    }
    sink = {
      ref = {
        kind        = "InMemoryChannel"
        api_version = "messaging.knative.dev/v1"
        name        = "messages"
      }
    }
  }
}
