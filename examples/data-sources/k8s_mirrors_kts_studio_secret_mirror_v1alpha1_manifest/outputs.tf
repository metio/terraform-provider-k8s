output "manifests" {
  value = {
    "example" = data.k8s_mirrors_kts_studio_secret_mirror_v1alpha1_manifest.example.yaml
  }
}
