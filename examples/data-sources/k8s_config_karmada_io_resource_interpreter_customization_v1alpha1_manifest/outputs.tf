output "manifests" {
  value = {
    "example" = data.k8s_config_karmada_io_resource_interpreter_customization_v1alpha1_manifest.example.yaml
  }
}
