output "manifests" {
  value = {
    "example" = data.k8s_secretgenerator_mittwald_de_string_secret_v1alpha1_manifest.example.yaml
  }
}
