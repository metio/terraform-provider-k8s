data "k8s_generators_external_secrets_io_github_access_token_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
