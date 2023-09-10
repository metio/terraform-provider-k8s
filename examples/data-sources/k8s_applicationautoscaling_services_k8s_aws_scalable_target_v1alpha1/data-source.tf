data "k8s_applicationautoscaling_services_k8s_aws_scalable_target_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
