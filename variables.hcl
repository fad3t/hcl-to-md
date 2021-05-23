variable "cluster_name" {
  type        = "String"
  description = "Name of the cluster"
}

variable "cluster_version" {
  type        = "String"
  description = "Version of the cluster"
  default     = "1.20"
}
