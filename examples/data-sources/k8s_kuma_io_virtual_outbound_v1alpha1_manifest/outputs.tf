output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_virtual_outbound_v1alpha1_manifest.example.yaml
  }
}
