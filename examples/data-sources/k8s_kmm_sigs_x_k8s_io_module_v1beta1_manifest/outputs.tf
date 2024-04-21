output "manifests" {
  value = {
    "example" = data.k8s_kmm_sigs_x_k8s_io_module_v1beta1_manifest.example.yaml
  }
}
