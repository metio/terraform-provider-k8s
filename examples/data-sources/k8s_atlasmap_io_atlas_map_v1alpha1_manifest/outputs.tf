output "manifests" {
  value = {
    "example" = data.k8s_atlasmap_io_atlas_map_v1alpha1_manifest.example.yaml
  }
}
