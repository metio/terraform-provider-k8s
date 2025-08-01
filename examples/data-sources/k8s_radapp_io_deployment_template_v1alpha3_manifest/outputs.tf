output "manifests" {
  value = {
    "example" = data.k8s_radapp_io_deployment_template_v1alpha3_manifest.example.yaml
  }
}
