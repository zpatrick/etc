variable "aws_access_key" {
  description = "Access Key portion of an AWS IAM key"
}

variable "aws_secret_key" {
  description = "Secret Key portion of an AWS IAM key"
}

variable "aws_region" {
  description = "AWS Region to build resources"
  default     = "us-west-2"
}

variable "name" {
  description = "Name of the cluster"
}
