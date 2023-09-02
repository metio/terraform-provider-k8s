output "manifests" {
  value = {
    "example" = data.k8s_apicodegen_apimatic_io_api_matic_v1beta1_manifest.example.yaml
  }
}
