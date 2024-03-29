output "manifests" {
  value = {
    "example" = data.k8s_velero_io_data_download_v2alpha1_manifest.example.yaml
  }
}
