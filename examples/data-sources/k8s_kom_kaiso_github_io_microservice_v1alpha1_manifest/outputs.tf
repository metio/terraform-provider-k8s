output "manifests" {
  value = {
    "example" = data.k8s_kom_kaiso_github_io_microservice_v1alpha1_manifest.example.yaml
  }
}
