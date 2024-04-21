data "k8s_k8s_mariadb_com_grant_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
