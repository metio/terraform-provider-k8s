data "k8s_storage_k8s_io_csi_driver_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    attach_required        = true
    pod_info_on_mount      = true
    volume_lifecycle_modes = ["Ephemeral"]
  }
}
