output "manifests" {
  value = {
    "example" = data.k8s_argoproj_io_application_set_v1alpha1_manifest.example.yaml
  }
}
