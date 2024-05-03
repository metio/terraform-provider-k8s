output "manifests" {
  value = {
    "example" = data.k8s_source_toolkit_fluxcd_io_helm_chart_v1_manifest.example.yaml
  }
}
