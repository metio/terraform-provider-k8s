output "manifests" {
  value = {
    "example" = data.k8s_velero_io_data_upload_v2alpha1_manifest.example.yaml
  }
}
