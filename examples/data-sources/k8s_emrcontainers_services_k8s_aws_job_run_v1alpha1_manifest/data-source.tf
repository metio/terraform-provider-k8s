data "k8s_emrcontainers_services_k8s_aws_job_run_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
