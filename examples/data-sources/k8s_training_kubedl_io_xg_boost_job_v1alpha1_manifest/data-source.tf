data "k8s_training_kubedl_io_xg_boost_job_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
