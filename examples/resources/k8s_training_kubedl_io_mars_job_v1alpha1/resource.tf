resource "k8s_training_kubedl_io_mars_job_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
