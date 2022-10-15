resource "k8s_service_account_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_service_account_v1" "example" {
  metadata = {
    name = "test"
  }
  secrets = [
    {
      name = "some-secret-name"
    },
  ]
}
