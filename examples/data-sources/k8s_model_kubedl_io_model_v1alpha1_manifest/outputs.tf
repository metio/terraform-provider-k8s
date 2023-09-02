output "manifests" {
  value = {
    "example" = data.k8s_model_kubedl_io_model_v1alpha1_manifest.example.yaml
  }
}
