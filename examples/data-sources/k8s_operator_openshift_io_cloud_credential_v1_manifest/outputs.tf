output "manifests" {
  value = {
    "example" = data.k8s_operator_openshift_io_cloud_credential_v1_manifest.example.yaml
  }
}
