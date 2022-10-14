resource "k8s_fossul_io_backup_config_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_fossul_io_backup_config_v1" "example" {
  metadata = {
    name = "mariadb-sample"
  }
  spec = {
    container_platform       = "openshift"
    operator_controlled      = true
    auto_discovery           = false
    deployment_name          = "mariadb"
    deployment_type          = "DeploymentConfig"
    job_retention            = 50
    overwrite_pcv_on_restore = true
    pvc_deletion_timeout     = 300
    restore_to_new_pvc       = false
    snapshot_timeout         = 180
    storage_plugin           = "csi.so"
    app_plugin               = "mariadb.so"
    policies = [
      {
        policy          = "hourly"
        retentionNumber = 3
      },
      {
        policy          = "daily"
        retentionNumber = 10
      },
    ]
  }
}
