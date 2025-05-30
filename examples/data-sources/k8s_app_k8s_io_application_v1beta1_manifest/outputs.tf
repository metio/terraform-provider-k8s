output "manifests" {
  value = {
    "example" = data.k8s_app_k8s_io_application_v1beta1_manifest.example.yaml
  }
}
