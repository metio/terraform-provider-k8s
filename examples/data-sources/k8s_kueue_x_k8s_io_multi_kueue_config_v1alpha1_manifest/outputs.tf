output "manifests" {
  value = {
    "example" = data.k8s_kueue_x_k8s_io_multi_kueue_config_v1alpha1_manifest.example.yaml
  }
}
