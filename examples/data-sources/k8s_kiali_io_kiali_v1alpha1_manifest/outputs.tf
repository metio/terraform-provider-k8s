output "manifests" {
  value = {
    "example" = data.k8s_kiali_io_kiali_v1alpha1_manifest.example.yaml
  }
}
