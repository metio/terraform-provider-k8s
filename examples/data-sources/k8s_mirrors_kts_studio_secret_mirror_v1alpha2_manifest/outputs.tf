output "manifests" {
  value = {
    "example" = data.k8s_mirrors_kts_studio_secret_mirror_v1alpha2_manifest.example.yaml
  }
}
