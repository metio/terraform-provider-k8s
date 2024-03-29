output "manifests" {
  value = {
    "example" = data.k8s_config_openshift_io_feature_gate_v1_manifest.example.yaml
  }
}
