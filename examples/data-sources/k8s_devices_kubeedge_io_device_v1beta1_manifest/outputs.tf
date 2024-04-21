output "manifests" {
  value = {
    "example" = data.k8s_devices_kubeedge_io_device_v1beta1_manifest.example.yaml
  }
}
