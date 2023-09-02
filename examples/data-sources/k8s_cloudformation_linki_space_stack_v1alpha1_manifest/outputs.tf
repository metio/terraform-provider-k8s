output "manifests" {
  value = {
    "example" = data.k8s_cloudformation_linki_space_stack_v1alpha1_manifest.example.yaml
  }
}
