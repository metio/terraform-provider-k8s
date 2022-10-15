resource "k8s_storage_k8s_io_csi_driver_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}

resource "k8s_storage_k8s_io_csi_driver_v1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    attach_required        = true
    pod_info_on_mount      = true
    volume_lifecycle_modes = ["Ephemeral"]
  }
}
