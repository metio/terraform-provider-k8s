output "manifests" {
  value = {
    "example" = data.k8s_operator_cryostat_io_cryostat_v1beta2_manifest.example.yaml
  }
}
