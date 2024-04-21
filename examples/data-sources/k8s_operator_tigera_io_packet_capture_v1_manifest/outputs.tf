output "manifests" {
  value = {
    "example" = data.k8s_operator_tigera_io_packet_capture_v1_manifest.example.yaml
  }
}
