output "manifests" {
  value = {
    "example" = data.k8s_operator_openshift_io_service_ca_v1_manifest.example.yaml
  }
}
