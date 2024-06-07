data "k8s_forklift_konveyor_io_ovirt_volume_populator_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    disk_id            = "some-id"
    engine_secret_name = "some-secret"
    engine_url         = "example.com"
  }
}
