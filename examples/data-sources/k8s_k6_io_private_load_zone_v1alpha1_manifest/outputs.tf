output "manifests" {
  value = {
    "example" = data.k8s_k6_io_private_load_zone_v1alpha1_manifest.example.yaml
  }
}
