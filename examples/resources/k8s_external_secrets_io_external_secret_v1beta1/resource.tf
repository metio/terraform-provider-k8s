resource "k8s_external_secrets_io_external_secret_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
