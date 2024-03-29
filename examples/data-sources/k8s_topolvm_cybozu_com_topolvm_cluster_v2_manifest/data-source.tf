data "k8s_topolvm_cybozu_com_topolvm_cluster_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
