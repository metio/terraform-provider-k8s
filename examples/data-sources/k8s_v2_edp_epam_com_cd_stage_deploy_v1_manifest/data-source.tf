data "k8s_v2_edp_epam_com_cd_stage_deploy_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
