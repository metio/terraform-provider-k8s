data "k8s_pipes_services_k8s_aws_pipe_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
