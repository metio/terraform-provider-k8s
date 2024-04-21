output "manifests" {
  value = {
    "example" = data.k8s_operator_marin3r_3scale_net_discovery_service_v1alpha1_manifest.example.yaml
  }
}
