output "manifests" {
  value = {
    "example" = data.k8s_infra_contrib_fluxcd_io_terraform_v1alpha2_manifest.example.yaml
  }
}
