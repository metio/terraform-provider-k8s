output "manifests" {
  value = {
    "example" = data.k8s_imageregistry_operator_openshift_io_config_v1_manifest.example.yaml
  }
}
