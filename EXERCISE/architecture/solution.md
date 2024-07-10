# SOLUTION

## 1. Bottlenecks and Risks

1. Hard to scale: this architecture is not designed for easy scaling. The are some workarounds for implementing autoscaling while still using VMs, but I think a much more effective solution would be to move to containerized environment where we can implement container orchestration tools for auto-scaling.

2. Single point of failure: all components of the application are running on standalone VMs. If any single one of these fails, the whole system goes down.

3. Bottlenecks: 
- the Load Balancer on a standalone VM can become a bottleneck under high traffic conditions. 
- the Document Processor on a single VM limits its ability to handle concurrent requests, causing potential performance bottlenecks.
- NFS volume can be a bottleneck as it might need to handle a large number of I/O operations from the document processor.

4. Redundancy and availability: we have data redundancy for the images with the RAID 1 solution. This protects us from disk failure, but not from server failure. If the VM goes down, we might have the images backed up but there would be no way to access them. Also lacks geographic redundancy.

### 2. New design

I would adopt a cloud native & microservice approach where we can leverage all the benefits of the cloud and orchestration tools like Kubernetes.

We would use Kubernetes as a container orchestrator, especially because of its extensive ecosystem.

We would deploy to a cloud provider like AWS to ensure scalability, reliability, and ease of management.

We would use AWS's Kubernetes service (EKS). AWS would handle the Kubernetes control plane which would reduce the operational burden on our team.

To acces the cluster we'd use AWS Load Balancer Controller. AWS Load Balancer Controller acts as our ingress controller but it also creates and manages AWS Elastic Load Balancers for the cluster. When a Kubernetes Service of type LoadBalancer is created, the AWS Load Balancer controller creates a Network Load Balancer (NLB). And when a Kubernetes Ingress object is created, the AWS Load Balancer Controller creates an Application Load Balancer (ALB). These balance traffic at Layer 4 and Layer 7 of the OSI model respectevely.

The recommended Load Balancer type for HTTP/HTTPS workloads is ALB. So we would use Ingress objects instead of Services of type LoadBalancer. The created ALB is controlled and configured by the Ingress object and routes HTTP or HTTPS traffic to different Pods within the cluster. This would give us the flexibility to change the application traffic routing as we please.

In terms of image storage Amazon S3: Stores images, providing high availability, durability, and scalability. S3's lifecycle policies can manage the retention and deletion of images over time.

With this setup we can:
1. In terms of scalability: we can take full advantage of both K8s and AWSs auto scaling solutions. Horizontal pod autoscaler for the Document Processor deployment. And AWS's Autoscaling Groups for automatic deployment of new EC2 worker nodes when a predefined CPU or RAM threshold has been reached in the running nodes.
2. In terms of single point of failure: Using EKS with multiple replicas and an ALB eliminates the single point of failure. The architecture is designed to be automatically scalable in every stage. Except for some data center wide issue... for this case we would need to set up a disaster recovery environment in another region.
3. In terms of bottlenecks: we don't have any bottlenecks anymore because all of our components scale automatically.
4. In terms of redundancy and availability: leveraging the power of S3, we can ensure.... S3 provides eleven 9's of durability and replicates data across multiple Availability Zones (AZs). EFS also provides multi-AZ redundancy.

### Further improvements

1. Add ElastiCache as a caching solution: store metadata such as document IDs, upload dates, and owner information in ElastiCache. This allows quick access to metadata without accessing the primary storage (S3) repeatedly.

2. Use Lambda for asynchronous processing: instead of handling document processing synchronously within the main application a Lambda functions can be triggered to do this. This allows for offloading tasks like image processing to Lambda, which can scale automatically based on incoming workload.

3. GitOps with ArgoCD: Implement GitOps methodology using ArgoCD. The Git repository will serve as the single source of truth for the desired state of the system, ensuring that any changes in the repository are automatically reflected in the environment.

4. Service Mesh with Istio: We could use Istio for traffic management, security, and observability. Istio provides advanced traffic management capabilities like A/B testing, canary releases, and fault injection. We could add Flagger to this for automated canary deployments

4. Monitoring and logging: Set up comprehensive monitoring and logging using Prometheus, Grafana, and Loki. This will provide insights into application performance and system health.


### Cons
In terms of cons for this whole setup we could say one is the complexity of Kubernetes and adopting GitOps. We would need to abstract these complexities away as much as possible from teams who don't need to know about it.

Another con would be costs and vendor-locking. We'd need to try to remain as provider agnostic as possible, but we would probably never be 100% free.


## 3. Action Plan

### 1. Planning and Communication

Design: Develop a detailed architecture plan, including diagrams and component descriptions.
Proof of Concept: Implement a small-scale proof of concept to validate the new architecture.

We would need to schedule meetings with development, operations, and customer service teams to communicate the new architecture and its benefits. Gather feedback and identify any specific requirements or concerns from each team. 

We'd develop a detailed project plan with timelines, milestones, and responsibilities. Here we'd identify potential risks and set up mitigation strategies.


### 2. Infrastructure Setup

Set up the AWS environment, including IAM roles, VPC, subnets, and security groups.
Create an EKS cluster and configure node groups.

deploy the document processor application.

CI/CD Pipeline:

Implement a CI/CD pipeline for automated deployment of the document processor application to EKS.
Integrate code repository (e.g., GitHub) with the CI/CD pipeline.

Configure S3 buckets.

### 3. Migration and Testing

Application Refactoring: Modify the document processor application to be containerized (Docker) and compatible with Kubernetes.
Update the application to use S3 for image storage and EFS for persistent storage.

Migrate existing documents from NFS to S3.
Set up data replication and backup strategies.

Testing:
Deploy the application to a staging environment in EKS.
Conduct thorough testing, including load testing, to ensure the application performs well under expected loads.

### 4. Rollout

Gradual Rollout:

Implement a phased rollout strategy, starting with a small subset of users.
Monitor the application and collect metrics to ensure stability and performance.

Full Deployment:
Once the application is validated in production, gradually increase the user base until the old system can be fully deprecated.

We could use a 

### 5. Monitoring and Optimization

Set up monitoring and alerting using Prometheus, and Grafana.
Continuously monitor application performance and system health.

Regularly review and optimize resource utilization. Implement cost-saving measures where possible.