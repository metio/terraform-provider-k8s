resource "k8s_couchbase_com_couchbase_backup_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    strategy = "full_incremental"
  }
}

resource "k8s_couchbase_com_couchbase_backup_v2" "example1" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    strategy = "full_incremental"
    auto_scaling = {
      limit = 123
    }
  }
}

resource "k8s_couchbase_com_couchbase_backup_v2" "example2" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    strategy = "full_incremental"
    auto_scaling = {
      limit = "150Ki"
    }
  }
}
