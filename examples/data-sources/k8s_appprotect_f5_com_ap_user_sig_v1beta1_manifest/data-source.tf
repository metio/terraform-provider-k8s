data "k8s_appprotect_f5_com_ap_user_sig_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
