data "k8s_operator_knative_dev_knative_eventing_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
