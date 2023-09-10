data "k8s_fossul_io_restore_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    deployment_name = "mariadb"
    policy          = "daily"
    workflow_id     = "xxxx"
  }
}
