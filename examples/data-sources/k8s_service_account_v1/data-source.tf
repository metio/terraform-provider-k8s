data "k8s_service_account_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
