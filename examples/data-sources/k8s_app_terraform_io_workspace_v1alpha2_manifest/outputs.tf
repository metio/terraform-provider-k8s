output "manifests" {
  value = {
    "example" = data.k8s_app_terraform_io_workspace_v1alpha2_manifest.example.yaml
  }
}
