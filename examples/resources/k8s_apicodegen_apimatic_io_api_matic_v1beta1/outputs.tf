output "resources" {
  value = {
    "minimal" = k8s_apicodegen_apimatic_io_api_matic_v1beta1.minimal.yaml
    "example" = k8s_apicodegen_apimatic_io_api_matic_v1beta1.example.yaml
  }
}
