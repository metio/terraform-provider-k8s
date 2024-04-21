output "manifests" {
  value = {
    "example" = data.k8s_kuadrant_io_dns_record_v1alpha1_manifest.example.yaml
  }
}
