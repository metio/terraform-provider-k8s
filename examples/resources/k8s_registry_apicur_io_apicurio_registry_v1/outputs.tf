output "resources" {
  value = {
    "minimal" = k8s_registry_apicur_io_apicurio_registry_v1.minimal.yaml
    "example" = k8s_registry_apicur_io_apicurio_registry_v1.example.yaml
  }
}
