output "manifests" {
  value = {
    "example" = data.k8s_kuadrant_io_managed_zone_v1alpha1_manifest.example.yaml
  }
}
