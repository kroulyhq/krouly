# Software Design Document: Web Data Extraction as a Service

This document represents the problem and provides a scalabe, flexible, and user-friendly solution for data extraction as a services. By using tools such as Terraform, Kubernetes, Go and a command line interface, a robust and efficinent infrastructure can be offered to meet all data extraction needs.

## Problem definition

The exponential growth of data on the web has created an increasing demand for efficient tools to extract and process relevant information. Many organizations and individuals require quick and reliable access to web data to feed their applications and make informed decisions. However, developing and maintaining infrastructure and developing data extraction tools can be complex and costly.

## Analysis

### Customer needs

- Build and deploy web data extraction workflows in minutes not months
- Access to web data reliably and effciently
- Ability to customize and control the data extraction according to their needs.
- Scalability to handle large volumes of data
- Flexibility in choosing infrastructure (cloud or on-premise)

### Current limitations

- Lack of robust and user-friendly tools for data extraction
- Complexity in configuring and managing infrastructure
- Risks of legal compliance breaches by not adhering to website's terms of service

## Architecture

### Main components

1. Sources: Websites and other data sources from which information will be extracted
2. Connectors: Modules that establish connection with sources and manage the extraction process
3. Connections: Configuration for data sources access
4. Destinations: Places where extracted data will be stored, such as database, data warehouse or file system
5. Playbooks: Steps for the workflow to follow in order to get to the source before extracting data

### Tools

- Terraform: Infrastructure management, allowing efficient and consistent deployment and maintenance of infrastructure
- Kubernetes: Container orchestration, providing scalability and high availability for services
- Go: Develop system components with a highly performance and efficient language
- Command line interface: Create and manage data extraction workflows
- Dashboard: Graphical interface to visualize and controlling workflow's status and activities

## Solution

Open sourced cloud based and on-promise ELT

### Connectors

- Audio analysis
- Lead generation
- Jobs boards
