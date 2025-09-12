output "manifests" {
  value = {
    "example" = data.k8s_fluxcd_controlplane_io_flux_report_v1_manifest.example.yaml
  }
}
