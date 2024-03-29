output "manifests" {
  value = {
    "example" = data.k8s_opentelemetry_io_open_telemetry_collector_v1beta1_manifest.example.yaml
  }
}
