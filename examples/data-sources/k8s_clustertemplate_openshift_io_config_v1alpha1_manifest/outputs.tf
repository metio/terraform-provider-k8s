output "manifests" {
  value = {
    "example" = data.k8s_clustertemplate_openshift_io_config_v1alpha1_manifest.example.yaml
  }
}
