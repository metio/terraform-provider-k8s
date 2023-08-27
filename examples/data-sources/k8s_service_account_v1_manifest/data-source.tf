data "k8s_service_account_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  secrets = [
    {
      name = "some-secret-name"
    },
  ]
}
