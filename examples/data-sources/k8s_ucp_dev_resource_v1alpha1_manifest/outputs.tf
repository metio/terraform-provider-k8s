output "manifests" {
  value = {
    "example" = data.k8s_ucp_dev_resource_v1alpha1_manifest.example.yaml
  }
}
