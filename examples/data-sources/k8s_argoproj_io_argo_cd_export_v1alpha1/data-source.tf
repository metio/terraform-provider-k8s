data "k8s_argoproj_io_argo_cd_export_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}