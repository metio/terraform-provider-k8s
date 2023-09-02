output "manifests" {
  value = {
    "example" = data.k8s_cert_manager_io_certificate_v1_manifest.example.yaml
  }
}
