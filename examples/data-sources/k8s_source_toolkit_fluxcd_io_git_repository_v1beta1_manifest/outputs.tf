output "manifests" {
  value = {
    "example" = data.k8s_source_toolkit_fluxcd_io_git_repository_v1beta1_manifest.example.yaml
  }
}
