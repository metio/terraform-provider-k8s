output "manifests" {
  value = {
    "example" = data.k8s_serving_kubedl_io_inference_v1alpha1_manifest.example.yaml
  }
}
