output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_zone_ingress_v1alpha1_manifest.example.yaml
  }
}
