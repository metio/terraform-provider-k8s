resource "k8s_mattermost_com_cluster_installation_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    ingress_name = "some-name"
  }
}
