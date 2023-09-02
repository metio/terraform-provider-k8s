data "k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
