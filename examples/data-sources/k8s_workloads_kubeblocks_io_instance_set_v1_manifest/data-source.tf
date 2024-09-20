data "k8s_workloads_kubeblocks_io_instance_set_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
