output "resources" {
  value = {
    "minimal" = k8s_apps_clusternet_io_helm_release_v1alpha1.minimal.yaml
  }
}
