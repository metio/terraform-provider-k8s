output "manifests" {
  value = {
    "example" = data.k8s_forklift_konveyor_io_network_map_v1beta1_manifest.example.yaml
  }
}
