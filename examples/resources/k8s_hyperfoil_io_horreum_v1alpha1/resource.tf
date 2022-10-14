resource "k8s_hyperfoil_io_horreum_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_hyperfoil_io_horreum_v1alpha1" "example" {
  metadata = {
    name = "horreum"
  }
  spec = {
    node_host = "127.0.0.1"
  }
}
