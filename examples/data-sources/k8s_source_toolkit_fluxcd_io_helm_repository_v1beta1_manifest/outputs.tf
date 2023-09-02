output "manifests" {
  value = {
    "example" = data.k8s_source_toolkit_fluxcd_io_helm_repository_v1beta1_manifest.example.yaml
  }
}
