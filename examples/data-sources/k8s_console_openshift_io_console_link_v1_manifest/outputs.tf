output "manifests" {
  value = {
    "example" = data.k8s_console_openshift_io_console_link_v1_manifest.example.yaml
  }
}
