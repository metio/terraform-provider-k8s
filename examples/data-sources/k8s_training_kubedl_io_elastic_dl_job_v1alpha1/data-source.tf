data "k8s_training_kubedl_io_elastic_dl_job_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}