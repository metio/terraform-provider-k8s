data "k8s_forklift_konveyor_io_forklift_controller_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
