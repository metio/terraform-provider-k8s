resource "k8s_scheduling_volcano_sh_pod_group_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
