output "manifests" {
  value = {
    "example" = data.k8s_networking_gke_io_managed_certificate_v1_manifest.example.yaml
  }
}
