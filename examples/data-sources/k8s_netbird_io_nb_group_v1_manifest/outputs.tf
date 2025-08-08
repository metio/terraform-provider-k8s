output "manifests" {
  value = {
    "example" = data.k8s_netbird_io_nb_group_v1_manifest.example.yaml
  }
}
