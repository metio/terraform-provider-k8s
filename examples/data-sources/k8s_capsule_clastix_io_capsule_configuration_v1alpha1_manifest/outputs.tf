output "manifests" {
  value = {
    "example" = data.k8s_capsule_clastix_io_capsule_configuration_v1alpha1_manifest.example.yaml
  }
}
