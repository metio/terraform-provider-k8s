output "manifests" {
  value = {
    "example" = data.k8s_forklift_konveyor_io_storage_map_v1beta1_manifest.example.yaml
  }
}
