output "manifests" {
  value = {
    "example" = data.k8s_dex_gpu_ninja_com_dex_identity_provider_v1alpha1_manifest.example.yaml
  }
}
