output "manifests" {
  value = {
    "example" = data.k8s_telemetry_istio_io_telemetry_v1_manifest.example.yaml
  }
}
