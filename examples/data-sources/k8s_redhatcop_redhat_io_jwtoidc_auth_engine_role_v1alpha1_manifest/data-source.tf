data "k8s_redhatcop_redhat_io_jwtoidc_auth_engine_role_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
