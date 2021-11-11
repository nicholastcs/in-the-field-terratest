variable "name" {
  type        = string
  description = "Name of VM. Random string will generated if it is empty."
  default     = ""
}

variable "environment" {
  type        = string
  description = "Environment tag."
  default     = "dev"
  validation {
    error_message = "Variable var.environment must be either prod, uat or dev."
    condition     = contains(["prod", "uat", "dev"], var.environment)
  }
}

variable "cpu_cores" {
  type        = number
  description = "Number of CPU cores."
  default     = 2
  validation {
    error_message = "Variable var.cpu_cores must be greater than 0."
    condition     = var.cpu_cores > 0
  }
}

variable "ram" {
  type        = string
  description = "Number of CPU cores."
  default     = "256"
  validation {
    error_message = "Variable var.ram must be either 256, 512, 1024 or 2048."
    condition     = contains(["256", "512", "1024", "2048"], var.ram)
  }
}
