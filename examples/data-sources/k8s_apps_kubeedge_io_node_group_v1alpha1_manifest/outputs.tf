output "manifests" {
  value = {
    "example" = data.k8s_apps_kubeedge_io_node_group_v1alpha1_manifest.example.yaml
  }
}
