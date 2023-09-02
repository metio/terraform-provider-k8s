output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_health_check_v1alpha1_manifest.example.yaml
  }
}
