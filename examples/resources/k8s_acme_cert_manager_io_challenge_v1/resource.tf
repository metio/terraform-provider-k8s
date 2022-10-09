resource "k8s_acme_cert_manager_io_challenge_v1" "big" {
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
    authorization_url = "https://example.com/auth"
    dns_name          = "some-name"
    url               = "https://example.com"
    wildcard          = true
    token             = "123abc"
    issuer_ref = {
      group = "first"
      kind  = "dns"
      name  = "acme"
    }
    type = "DNS-01"
    solver = {
      selector = {
        match_labels = {
          "someLabel" = "someValue"
        }
      }
      dns01 = {
        acme_dns = {
          host = "acme-company.com"
          account_secret_ref = {
            name = "acme-credentials"
            key  = "nested-key"
          }
        }
      }
    }
    key = "12345"
  }
}

resource "k8s_acme_cert_manager_io_challenge_v1" "small" {
  metadata = {
    name = "test"
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
