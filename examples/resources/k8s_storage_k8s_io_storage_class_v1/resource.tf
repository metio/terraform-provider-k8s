resource "k8s_storage_k8s_io_storage_class_v1" "minimal" {
  metadata = {
    name = "test"
  }
  provisioner = "kubernetes.io/gce-pd"
}

resource "k8s_storage_k8s_io_storage_class_v1" "example" {
  metadata = {
    name = "test"
  }
  provisioner    = "kubernetes.io/gce-pd"
  reclaim_policy = "Retain"
  parameters = {
    type = "pd-standard"
  }
  mount_options = ["file_mode=0700", "dir_mode=0777", "mfsymlinks", "uid=1000", "gid=1000", "nobrl", "cache=none"]
}
