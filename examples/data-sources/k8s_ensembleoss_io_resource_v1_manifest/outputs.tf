output "manifests" {
  value = {
    "example" = data.k8s_ensembleoss_io_resource_v1_manifest.example.yaml
  }
}
