resource "k8s_metal3_io_preprovisioning_image_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
