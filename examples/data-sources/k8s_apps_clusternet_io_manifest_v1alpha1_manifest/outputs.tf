output "manifests" {
  value = {
    "example" = data.k8s_apps_clusternet_io_manifest_v1alpha1_manifest.example.yaml
  }
}
