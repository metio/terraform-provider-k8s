data "k8s_druid_apache_org_druid_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    common_runtime_properties = "some-property"
    nodes = {
      druid_port             = 12345
      node_config_mount_path = "/mnt/path"
      node_type              = "overlord"
      runtime_properties     = "some-property"
    }
  }
}
