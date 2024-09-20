output "manifests" {
  value = {
    "example" = data.k8s_apps_kubeblocks_io_component_version_v1_manifest.example.yaml
  }
}
