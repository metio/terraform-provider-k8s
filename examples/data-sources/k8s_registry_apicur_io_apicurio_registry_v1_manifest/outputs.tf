output "manifests" {
  value = {
    "example" = data.k8s_registry_apicur_io_apicurio_registry_v1_manifest.example.yaml
  }
}
