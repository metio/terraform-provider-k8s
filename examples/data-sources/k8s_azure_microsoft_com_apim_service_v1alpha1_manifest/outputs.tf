output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_apim_service_v1alpha1_manifest.example.yaml
  }
}
