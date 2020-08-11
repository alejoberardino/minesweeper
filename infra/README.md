# Minesweeper Infra

This is a _very_ quick sample tf to create an ec2 instance and an ELB to get a DNS address, the idea is to just run everything we do in the local env, but in that instance. If I had time, I'd do
- Static website hosted in S3, Served through Cloudfront 
- For something like a quick prototype like this, probably run the services in ECS using Fargate

## Prerequisites

- Install & Configure AWS CLI
- Install terraform

## Running terraform 

Run with a command, like this:

```sh
terraform apply -var 'key_name={your_aws_key_name}' \
   -var 'public_key_path={location_of_your_key_in_your_local_machine}'
```