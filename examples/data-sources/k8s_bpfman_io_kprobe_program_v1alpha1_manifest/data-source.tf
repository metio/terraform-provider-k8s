data "k8s_bpfman_io_kprobe_program_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }

  spec = {
    bpffunctionname = "some-value"
    bytecode        = {}
    func_name       = "some-value"
    nodeselector    = {}
  }
}
