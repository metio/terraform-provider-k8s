output "manifests" {
  value = {
    "example" = data.k8s_parameters_kubeblocks_io_component_parameter_v1alpha1_manifest.example.yaml
  }
}
