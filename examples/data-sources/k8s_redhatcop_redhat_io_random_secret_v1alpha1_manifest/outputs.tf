output "manifests" {
  value = {
    "example" = data.k8s_redhatcop_redhat_io_random_secret_v1alpha1_manifest.example.yaml
  }
}
