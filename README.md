# In The Field: Terratest

Often, Terraform Modules (TFM) will used by software developers in a self-serve basis. Therefore, an end-to-end test suite can help to validate the TFM before deploying it to the production. 

This helps on governance, security and provide a guard-rail for developer on cloud landing-zone infrastructure provisioning.

Terratest is the de-facto test suite that will bring the best of Go programming language to carry out assertion and acceptance tests. 

It can invoke `terraform` commands through Go, therefore, the output can be used to carry out various tests like REST calls, ICMP, string assertion, expect fail and etc. Test capabilities are not bounded by Terraform limits but Go programming language.

