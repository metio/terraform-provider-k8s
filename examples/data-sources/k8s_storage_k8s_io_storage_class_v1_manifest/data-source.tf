data "k8s_storage_k8s_io_storage_class_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  k8s_provisioner = "kubernetes.io/gce-pd"
  reclaim_policy  = "Retain"
  parameters = {
    type = "pd-standard"
  }
  mount_options = ["file_mode=0700", "dir_mode=0777", "mfsymlinks", "uid=1000", "gid=1000", "nobrl", "cache=none"]
}
