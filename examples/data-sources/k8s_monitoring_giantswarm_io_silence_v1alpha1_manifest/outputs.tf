output "manifests" {
  value = {
    "example" = data.k8s_monitoring_giantswarm_io_silence_v1alpha1_manifest.example.yaml
  }
}
