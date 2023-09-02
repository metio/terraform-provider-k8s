resource "k8s_operator_knative_dev_knative_serving_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
