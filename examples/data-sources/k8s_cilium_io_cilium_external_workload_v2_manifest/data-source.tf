data "k8s_cilium_io_cilium_external_workload_v2_manifest" "example" {
  metadata = {
    name = "some-name"
  }
}
