output "manifests" {
  value = {
    "example" = data.k8s_integration_rock8s_com_deferred_resource_v1beta1_manifest.example.yaml
  }
}
