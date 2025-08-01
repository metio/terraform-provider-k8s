output "manifests" {
  value = {
    "example" = data.k8s_workspace_devfile_io_dev_workspace_template_v1alpha1_manifest.example.yaml
  }
}
