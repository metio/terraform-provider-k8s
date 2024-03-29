data "k8s_clustertemplate_openshift_io_cluster_template_instance_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
