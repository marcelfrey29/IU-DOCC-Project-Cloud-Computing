# Idea for the Project

The idea is to build a simple thee-tier web application with a frontend-service, backend-service and a database.

## High-level Project Plan

1. Define requirements for the application
1. Develop a local version of the application
    - Host it locally (e.g. via [Docker](https://www.docker.com/) Containers)
1. Research, compare and select Service Models and Cloud Services (description, pros, cons, comparison, selection)
    - Maybe have a look at the _AWS Migration Strategies (6Rs)_
    - Evaluate, compare, and discuss different approaches
1. Define requirements for migration of the application
1. Migrate the application to the cloud (refactor and deploy)

## Idea for the Application: Travel Guide App

The applications allows users to create, edit, and view travel guides. 
Travel Guides contain activities and attractions for a specific location like a city.
Visitors of this location can use Travel Guides for inspiration and orientation.

The application doesn't support user accounts or user-specific data. 
Travel Guides, attractions and activities are public by default.
However, Travel Guides can be protected by a secret value that can be shared with other travellers for collaboration.

- A travel guide is for a specific location (e.g. a city)
- Inside the guide, certain attractions and activities can be listed
- Attractions and activities can contain details like location, expected duration, pricing information, etc.
- Creators of travel guides must define an editing secret (password) which is required for edits
    - Collaboration is possible by sharing the secret with other travellers
- Travel guides can be made private (The secret is required to view it)

## Technical Ideas

As there are no special requirements regarding the technologies, I'll explore some new ones in this project. 

- Web-App: [React](https://react.dev/) and [NextUI](https://nextui.org/)
- Backend: [Golang](https://go.dev/) and [Fiber](https://gofiber.io/)
- Database: [DynamoDB](https://aws.amazon.com/dynamodb/)
    - There is a local version [DynamoDB Local](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html) and other DynamoDB-compatible Databses like [ScyllaDB](https://www.scylladb.com/alternator/)
- Cloud Provider: [AWS](https://aws.amazon.com/)
- Cloud Deployment: Infrastructre as Code with [AWS CloudFormation](https://aws.amazon.com/cloudformation/)
