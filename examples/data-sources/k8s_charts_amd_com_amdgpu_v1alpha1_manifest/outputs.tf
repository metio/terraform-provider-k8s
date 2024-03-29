output "manifests" {
  value = {
    "example" = data.k8s_charts_amd_com_amdgpu_v1alpha1_manifest.example.yaml
  }
}
