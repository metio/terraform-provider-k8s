data "k8s_inference_kubedl_io_elastic_batch_job_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
