output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_blob_container_v1alpha2_manifest.example.yaml
  }
}
