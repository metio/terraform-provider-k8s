resource "k8s_cilium_io_cilium_external_workload_v2" "minimal" {
  metadata = {
    name = "test"
  }
}
