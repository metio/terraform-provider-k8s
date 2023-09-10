data "k8s_argoproj_io_argo_cd_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
