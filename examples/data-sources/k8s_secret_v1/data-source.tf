data "k8s_secret_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
