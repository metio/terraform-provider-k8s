resource "k8s_argoproj_io_argo_cd_export_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_argoproj_io_argo_cd_export_v1alpha1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    argocd = "argocd-sample"
  }
}
