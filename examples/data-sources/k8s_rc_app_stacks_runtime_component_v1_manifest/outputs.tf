output "manifests" {
  value = {
    "example" = data.k8s_rc_app_stacks_runtime_component_v1_manifest.example.yaml
  }
}
