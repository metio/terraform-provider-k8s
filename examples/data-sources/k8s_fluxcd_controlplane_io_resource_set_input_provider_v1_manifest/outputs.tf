output "manifests" {
  value = {
    "example" = data.k8s_fluxcd_controlplane_io_resource_set_input_provider_v1_manifest.example.yaml
  }
}
