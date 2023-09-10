data "k8s_acme_cert_manager_io_challenge_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    authorization_url = "https://example.com/auth"
    dns_name          = "some-name"
    url               = "https://example.com"
    token             = "123abc"
    issuer_ref = {
      group = "first"
      kind  = "dns"
      name  = "acme"
    }
    type = "HTTP-01"
    solver = {
      selector = {
        match_labels = {
          "someLabel" = "someValue"
        }
      }
    }
    key = "12345"
  }
}
