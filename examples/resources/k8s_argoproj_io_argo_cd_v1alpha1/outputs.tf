output "resources" {
  value = {
    "minimal" = k8s_argoproj_io_argo_cd_v1alpha1.minimal.yaml
    "example" = k8s_argoproj_io_argo_cd_v1alpha1.example.yaml
  }
}
