output "manifests" {
  value = {
    "example" = data.k8s_rc_app_stacks_runtime_operation_v1beta2_manifest.example.yaml
  }
}
