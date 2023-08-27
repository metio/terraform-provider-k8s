output "manifests" {
  value = {
    "example" = data.k8s_storage_k8s_io_csi_driver_v1_manifest.example.yaml
  }
}
