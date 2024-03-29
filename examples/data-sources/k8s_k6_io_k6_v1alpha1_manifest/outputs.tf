output "manifests" {
  value = {
    "example" = data.k8s_k6_io_k6_v1alpha1_manifest.example.yaml
  }
}
