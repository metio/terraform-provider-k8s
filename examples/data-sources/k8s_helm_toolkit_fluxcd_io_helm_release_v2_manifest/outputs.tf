output "manifests" {
  value = {
    "example" = data.k8s_helm_toolkit_fluxcd_io_helm_release_v2_manifest.example.yaml
  }
}
