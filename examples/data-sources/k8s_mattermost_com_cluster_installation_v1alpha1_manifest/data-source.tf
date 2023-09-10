data "k8s_mattermost_com_cluster_installation_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    ingress_name = "some-name"
  }
}
