output "manifests" {
  value = {
    "example" = data.k8s_helm_sigstore_dev_rekor_v1alpha1_manifest.example.yaml
  }
}
