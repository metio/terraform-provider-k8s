output "manifests" {
  value = {
    "example" = data.k8s_source_toolkit_fluxcd_io_oci_repository_v1beta2_manifest.example.yaml
  }
}
