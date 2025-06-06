output "manifests" {
  value = {
    "example" = data.k8s_parameters_kubeblocks_io_param_config_renderer_v1alpha1_manifest.example.yaml
  }
}
