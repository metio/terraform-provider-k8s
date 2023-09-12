output "manifests" {
  value = {
    "example" = data.k8s_boskos_k8s_io_drlc_object_v1_manifest.example.yaml
  }
}
