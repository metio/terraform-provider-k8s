output "manifests" {
  value = {
    "example" = data.k8s_core_openfeature_dev_feature_flag_configuration_v1alpha1_manifest.example.yaml
  }
}
