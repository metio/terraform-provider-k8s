output "manifests" {
  value = {
    "example" = data.k8s_nativestor_alauda_io_raw_device_v1_manifest.example.yaml
  }
}
