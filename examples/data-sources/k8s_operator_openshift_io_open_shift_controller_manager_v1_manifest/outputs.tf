output "manifests" {
  value = {
    "example" = data.k8s_operator_openshift_io_open_shift_controller_manager_v1_manifest.example.yaml
  }
}
