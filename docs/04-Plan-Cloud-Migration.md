# Plan Cloud Migration

Migrate the local version of the application to the AWS Cloud.

## Criteria

- Sclalability and Fit-to-Workload
- Availability
- Security
    - In Transit
    - At Rest
- Cost (Infrastructure)
- Operational Model
    - Capacity Planning
    - Maintenance / Patching
- Risk of Lock-In

> [!WARNING] 
> For Scalability and Fit-to-Workload we have to assume a usage scenario. 
>
> We assume that our application is heaviliy used before and after the vacation-season. 
> This means we expect four spikes over the year (Before Summer, After Summer, Before Winter, After Winter).
>
> Before users go to vaction, they get inspiration on what to do at their location. 
> While the load is higher in general, we will have many read-operations.
>
> Once people come back, they might create their own travel guide to inspire other travelers. 
> While the load is higher in general, we will have more write-operations. 

>[!WARNING]
> To make a fair price comparison, we must compare the solutions with the same level of availability and scalability.
> 
> For example, AWS Lambda is hightly available by default because it uses all Availability Zones (Re-architecture). 
> While a single server would be enough for the Rehost and Replatform solution, we will calculate the price with 3 instances and a load balancer to have the same level of availability.

>[!NOTE]
> Some general criteria apply to all migration strategies
> - Deployment Model (Public Cloud)
> - Utility Pricing / No upfront costs
> - Internet Connectivity / Outages
> - Data Deletion

## References 

- [AWS 6Rs](https://aws.amazon.com/blogs/enterprise-strategy/6-strategies-for-migrating-applications-to-the-cloud/)
    - **Rehosting**: Lift-and-shift
        - Move existing workloads as they are
        - No cloud optimization
    - **Replatforming**: Lift-thinker-and-shift
        - Core Architecture stays unchanged
        - A few cloud optimizations are performed
    - **Refactoring/Re-architecting**
        - Use cloud-native features

## Possible Migration Strategies

**Retire**, **Retain**, and **Repurchasing** are not allowed for this project. 😅

### Rehosting

- IaaS
- Use a single EC2 Virtual Machine, install Docker, and run the Compose Stack
    - _We continue with our own database container (DynamoDB Local is not for production, but there are DDB-compatible databases that could be used)_
    - General Prupose Instance with AWS Gravition (`t4g`), 4vCPUs, 16GB Memory
    - 128GB Storage (DB Data are stored on this instance)
    - Hourly EBS Snapshots (we assume that ­2GB of data change per hour: DB+Logs)
    - 500GB of data are transferred to the internet
- Scaling doesn't work out of the box because the database is part of the compose stack
- Data are stoed on the public-facing EC2 Instance

![Rehost Architecture Diagram](assets/migration-rehost.svg)

- **Sclalability and Fit-to-Workload**:
    - Scaling is not possible out of the box (DB must be shared and can't be part of the stack)
    - We have only a single instance all the time
- **Availability**: Single host is single point of failure, no redundancy 
- **Security**:
    - Database/Data on public facting instance (compromised host, compromised data)
    - **In Transit**: TLS Certificate can be added to nginx, e.g. via Let's Encrypt (Manual Work)
    - **At Rest**: Usage of encrypted EC2 Volume is possible
- **Cost (Infrastructure)**: 206,04 USD / Month
- **Operational Model**:
    - **Capacity Planning**: _Doesn't apply, not possible_
    - **Maintenance / Patching**: We need to keep the instance up-to-date
- **Risk of Lock-In**: None, we can launch the Docker Compose Stack on any Linux VM


