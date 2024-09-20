output "manifests" {
  value = {
    "example" = data.k8s_listeners_stackable_tech_pod_listeners_v1alpha1_manifest.example.yaml
  }
}
