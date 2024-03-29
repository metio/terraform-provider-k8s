data "k8s_beegfs_csi_netapp_com_beegfs_driver_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
