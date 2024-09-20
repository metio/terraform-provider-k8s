output "manifests" {
  value = {
    "example" = data.k8s_listeners_stackable_tech_listener_class_v1alpha1_manifest.example.yaml
  }
}
