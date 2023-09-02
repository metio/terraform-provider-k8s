output "manifests" {
  value = {
    "example" = data.k8s_aquasecurity_github_io_aqua_starboard_v1alpha1_manifest.example.yaml
  }
}
