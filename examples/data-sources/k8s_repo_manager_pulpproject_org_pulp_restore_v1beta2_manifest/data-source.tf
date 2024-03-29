data "k8s_repo_manager_pulpproject_org_pulp_restore_v1beta2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
