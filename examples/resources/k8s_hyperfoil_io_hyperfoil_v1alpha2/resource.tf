resource "k8s_hyperfoil_io_hyperfoil_v1alpha2" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_hyperfoil_io_hyperfoil_v1alpha2" "example" {
  metadata = {
    name = "hyperfoil"
  }
  spec = {
    version = "latest"
  }
}
