output "manifests" {
  value = {
    "example" = data.k8s_theketch_io_app_v1beta1_manifest.example.yaml
  }
}
