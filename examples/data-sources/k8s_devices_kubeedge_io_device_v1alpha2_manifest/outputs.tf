output "manifests" {
  value = {
    "example" = data.k8s_devices_kubeedge_io_device_v1alpha2_manifest.example.yaml
  }
}
