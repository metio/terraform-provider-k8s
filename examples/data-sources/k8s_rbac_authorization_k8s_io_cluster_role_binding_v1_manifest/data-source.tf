data "k8s_rbac_authorization_k8s_io_cluster_role_binding_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  role_ref = {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster-admin"
  }
}
