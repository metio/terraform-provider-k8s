output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_app_insights_api_key_v1alpha1_manifest.example.yaml
  }
}
