output "manifests" {
  value = {
    "example" = data.k8s_secrets_doppler_com_doppler_secret_v1alpha1_manifest.example.yaml
  }
}
