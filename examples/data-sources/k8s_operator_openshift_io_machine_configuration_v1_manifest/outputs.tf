output "manifests" {
  value = {
    "example" = data.k8s_operator_openshift_io_machine_configuration_v1_manifest.example.yaml
  }
}
