resource "k8s_acme_cert_manager_io_challenge_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
