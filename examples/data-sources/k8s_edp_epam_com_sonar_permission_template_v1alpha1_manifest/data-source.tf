data "k8s_edp_epam_com_sonar_permission_template_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
