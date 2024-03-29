data "k8s_kinesis_services_k8s_aws_stream_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
