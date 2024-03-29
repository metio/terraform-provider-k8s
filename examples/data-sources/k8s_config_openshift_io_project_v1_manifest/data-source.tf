data "k8s_config_openshift_io_project_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
