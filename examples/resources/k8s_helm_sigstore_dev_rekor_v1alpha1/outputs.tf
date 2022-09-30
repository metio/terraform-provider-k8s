output "resources" {
  value = {
    "minimal" = k8s_helm_sigstore_dev_rekor_v1alpha1.minimal.yaml
  }
}
