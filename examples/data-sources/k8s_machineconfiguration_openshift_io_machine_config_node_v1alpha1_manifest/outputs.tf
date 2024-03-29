output "manifests" {
  value = {
    "example" = data.k8s_machineconfiguration_openshift_io_machine_config_node_v1alpha1_manifest.example.yaml
  }
}
