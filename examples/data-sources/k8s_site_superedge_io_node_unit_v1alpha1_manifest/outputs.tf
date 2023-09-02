output "manifests" {
  value = {
    "example" = data.k8s_site_superedge_io_node_unit_v1alpha1_manifest.example.yaml
  }
}
