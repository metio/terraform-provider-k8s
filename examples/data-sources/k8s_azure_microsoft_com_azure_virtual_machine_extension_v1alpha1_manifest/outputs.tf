output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_azure_virtual_machine_extension_v1alpha1_manifest.example.yaml
  }
}
