resource "k8s_hyperfoil_io_hyperfoil_v1alpha2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
