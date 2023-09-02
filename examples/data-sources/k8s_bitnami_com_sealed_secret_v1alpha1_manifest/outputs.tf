output "manifests" {
  value = {
    "example" = data.k8s_bitnami_com_sealed_secret_v1alpha1_manifest.example.yaml
  }
}
