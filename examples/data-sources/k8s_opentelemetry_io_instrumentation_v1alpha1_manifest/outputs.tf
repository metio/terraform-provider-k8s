output "manifests" {
  value = {
    "example" = data.k8s_opentelemetry_io_instrumentation_v1alpha1_manifest.example.yaml
  }
}
