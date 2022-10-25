resource "k8s_operator_knative_dev_knative_eventing_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
