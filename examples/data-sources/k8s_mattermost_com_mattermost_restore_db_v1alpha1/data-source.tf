data "k8s_mattermost_com_mattermost_restore_db_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
