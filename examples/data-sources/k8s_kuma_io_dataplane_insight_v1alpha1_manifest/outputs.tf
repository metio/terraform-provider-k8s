output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_dataplane_insight_v1alpha1_manifest.example.yaml
  }
}
