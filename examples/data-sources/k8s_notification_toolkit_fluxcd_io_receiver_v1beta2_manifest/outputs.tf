output "manifests" {
  value = {
    "example" = data.k8s_notification_toolkit_fluxcd_io_receiver_v1beta2_manifest.example.yaml
  }
}
