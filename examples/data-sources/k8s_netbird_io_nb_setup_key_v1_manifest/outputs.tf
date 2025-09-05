output "manifests" {
  value = {
    "example" = data.k8s_netbird_io_nb_setup_key_v1_manifest.example.yaml
  }
}
