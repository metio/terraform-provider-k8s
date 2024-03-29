output "manifests" {
  value = {
    "example" = data.k8s_isindir_github_com_sops_secret_v1alpha1_manifest.example.yaml
  }
}
