output "manifests" {
  value = {
    "example" = data.k8s_secretgenerator_mittwald_de_ssh_key_pair_v1alpha1_manifest.example.yaml
  }
}
