output "manifests" {
  value = {
    "example" = data.k8s_kuadrant_io_dns_health_check_probe_v1alpha1_manifest.example.yaml
  }
}
