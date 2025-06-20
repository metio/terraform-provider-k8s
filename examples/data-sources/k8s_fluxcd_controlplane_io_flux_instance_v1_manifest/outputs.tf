output "manifests" {
  value = {
    "example" = data.k8s_fluxcd_controlplane_io_flux_instance_v1_manifest.example.yaml
  }
}
