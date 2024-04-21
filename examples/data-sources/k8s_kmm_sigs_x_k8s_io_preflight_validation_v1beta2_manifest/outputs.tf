output "manifests" {
  value = {
    "example" = data.k8s_kmm_sigs_x_k8s_io_preflight_validation_v1beta2_manifest.example.yaml
  }
}
