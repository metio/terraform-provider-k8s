output "manifests" {
  value = {
    "example" = data.k8s_addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha4_manifest.example.yaml
  }
}
