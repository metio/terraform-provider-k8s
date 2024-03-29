output "manifests" {
  value = {
    "example" = data.k8s_craftypath_github_io_sops_secret_v1alpha1_manifest.example.yaml
  }
}
