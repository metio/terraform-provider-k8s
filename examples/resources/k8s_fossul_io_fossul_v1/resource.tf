resource "k8s_fossul_io_fossul_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_fossul_io_fossul_v1" "example" {
  metadata = {
    name = "fossul-sample"
  }
  spec = {
    enable_routes          = false
    enable_virtualmachines = false
    container_platform     = "openshift"
  }
}
