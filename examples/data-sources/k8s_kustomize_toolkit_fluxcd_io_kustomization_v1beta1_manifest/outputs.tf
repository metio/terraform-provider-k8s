output "manifests" {
  value = {
    "example" = data.k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta1_manifest.example.yaml
  }
}
