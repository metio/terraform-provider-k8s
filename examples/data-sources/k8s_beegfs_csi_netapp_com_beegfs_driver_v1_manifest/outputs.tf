output "manifests" {
  value = {
    "example" = data.k8s_beegfs_csi_netapp_com_beegfs_driver_v1_manifest.example.yaml
  }
}
