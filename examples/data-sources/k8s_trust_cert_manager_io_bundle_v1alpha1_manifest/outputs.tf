output "manifests" {
  value = {
    "example" = data.k8s_trust_cert_manager_io_bundle_v1alpha1_manifest.example.yaml
  }
}
