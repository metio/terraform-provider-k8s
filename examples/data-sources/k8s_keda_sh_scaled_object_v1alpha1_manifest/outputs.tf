output "manifests" {
  value = {
    "example" = data.k8s_keda_sh_scaled_object_v1alpha1_manifest.example.yaml
  }
}
