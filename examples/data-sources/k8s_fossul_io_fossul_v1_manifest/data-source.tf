data "k8s_fossul_io_fossul_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    enable_routes          = false
    enable_virtualmachines = false
    container_platform     = "openshift"
  }
}
