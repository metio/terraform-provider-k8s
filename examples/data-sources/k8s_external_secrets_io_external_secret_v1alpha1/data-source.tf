data "k8s_external_secrets_io_external_secret_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}