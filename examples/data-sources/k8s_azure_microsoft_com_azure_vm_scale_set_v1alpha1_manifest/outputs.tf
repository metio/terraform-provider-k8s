output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_azure_vm_scale_set_v1alpha1_manifest.example.yaml
  }
}
