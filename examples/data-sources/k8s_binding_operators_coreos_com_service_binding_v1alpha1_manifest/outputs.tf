output "manifests" {
  value = {
    "example" = data.k8s_binding_operators_coreos_com_service_binding_v1alpha1_manifest.example.yaml
  }
}
