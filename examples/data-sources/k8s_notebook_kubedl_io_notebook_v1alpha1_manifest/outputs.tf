output "manifests" {
  value = {
    "example" = data.k8s_notebook_kubedl_io_notebook_v1alpha1_manifest.example.yaml
  }
}
