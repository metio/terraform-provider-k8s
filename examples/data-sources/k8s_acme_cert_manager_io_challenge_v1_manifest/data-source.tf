data "k8s_acme_cert_manager_io_challenge_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}