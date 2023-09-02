output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_traffic_trace_v1alpha1_manifest.example.yaml
  }
}
