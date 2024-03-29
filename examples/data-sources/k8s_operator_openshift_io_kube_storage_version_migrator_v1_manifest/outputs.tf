output "manifests" {
  value = {
    "example" = data.k8s_operator_openshift_io_kube_storage_version_migrator_v1_manifest.example.yaml
  }
}
