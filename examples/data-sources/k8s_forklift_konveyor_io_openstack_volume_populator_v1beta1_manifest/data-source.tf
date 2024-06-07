data "k8s_forklift_konveyor_io_openstack_volume_populator_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    identity_url = "example.com"
    image_id     = "some-image"
    secret_name  = "some-secret"
  }
}
