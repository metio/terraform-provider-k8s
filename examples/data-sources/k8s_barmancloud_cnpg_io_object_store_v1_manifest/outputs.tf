output "manifests" {
  value = {
    "example" = data.k8s_barmancloud_cnpg_io_object_store_v1_manifest.example.yaml
  }
}
