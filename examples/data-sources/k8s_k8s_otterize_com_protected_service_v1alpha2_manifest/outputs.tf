output "manifests" {
  value = {
    "example" = data.k8s_k8s_otterize_com_protected_service_v1alpha2_manifest.example.yaml
  }
}
