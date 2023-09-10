data "k8s_digitalis_io_vals_secret_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
