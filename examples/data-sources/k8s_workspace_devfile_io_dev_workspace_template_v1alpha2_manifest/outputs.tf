output "manifests" {
  value = {
    "example" = data.k8s_workspace_devfile_io_dev_workspace_template_v1alpha2_manifest.example.yaml
  }
}
