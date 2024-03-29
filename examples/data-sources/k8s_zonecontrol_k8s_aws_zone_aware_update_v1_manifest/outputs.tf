output "manifests" {
  value = {
    "example" = data.k8s_zonecontrol_k8s_aws_zone_aware_update_v1_manifest.example.yaml
  }
}
