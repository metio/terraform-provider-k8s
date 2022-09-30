output "resources" {
  value = {
    "minimal" = k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta2.minimal.yaml
  }
}
