
Atlas is an open-source database schema management and migration tool developed by Ariga. It helps engineers define, version, and evolve database schemas declaratively using code or configuration files. Atlas is designed to simplify database change workflows, making schema management more reliable and automated across development and production environments.

### Key facts

- **Developer:** Ariga
    
- **Initial release:** 2021
    
- **Languages supported:** SQL, HCL, and other schema definitions
    
- **Primary use:** Database schema management and migration automation
    
- **License:** Apache 2.0 open source
    

### Schema management and workflow

Atlas models a database schema as a declarative specification written in SQL or in its domain-specific HCL format. Developers can use `atlas migrate` commands to generate, inspect, and apply migration files safely. Atlas automatically computes diffs between the current database state and desired schema, generating migration steps accordingly. Its declarative approach allows teams to maintain schema definitions in version control alongside application code.

### Integration and automation

Atlas integrates with common development tools and CI/CD systems to enable continuous delivery of schema changes. It supports multiple relational database engines such as PostgreSQL, MySQL, MariaDB, SQLite, and CockroachDB. Integration with tools like Terraform and popular frameworks enables consistent schema management across environments, from local development to cloud deployments.

### Features and ecosystem

Atlas offers features such as migration linting, dependency visualization, and drift detection. It includes a cloud-hosted platform—Atlas Cloud—for collaborative management, visual schema inspection, and migration approvals. Developers can embed Atlas into their deployment pipelines or run it via CLI or API, ensuring schema evolution remains predictable, reversible, and testable.

### Adoption and community

Used by modern engineering teams seeking infrastructure-as-code principles for databases, Atlas has gained traction among open-source and enterprise users. Its focus on safety, automation, and developer ergonomics positions it as a leading solution for managing database schema lifecycle in DevOps workflows.