output "manifests" {
  value = {
    "example" = data.k8s_argoproj_io_argo_cd_v1alpha1_manifest.example.yaml
  }
}
