data "k8s_machine_deletion_remediation_medik8s_io_machine_deletion_remediation_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
