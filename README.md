# Snippet

Platform to share code snippets easily with others. (Gist or Pastebin Clone)

## Table of Contents 

- [Quick Start](#quick-start)
- [Features](#features)
- [Architucture](#architucture)
- [License](#license)

<a id="quick-start"></a>
## Quick Start

Snippet uses docker for easy development, docker and docker-compose are pre-requistes. Follow the following steps once the repository is cloned locally:

1. Create a `.env` file using example.env as an base
2. Run `docker-compose up --build`

Note: In order to test out if this is working as expected you can `curl http://localhost:8080/ping`, where the result should be `pong`

<a id="features"></a>
## Features

The core idea for the platform, is to allow users to create repos with up to 5 code snippets. Users can then share the snippets to other people.

- [X] Authentication System: Users are able register/login/logout
- [X] Built-in editor: Snippet uses Monaco editor, which is the same one used in VS Code, enabling powerful editing capabilities
- [X] Code File Storage: Snippet uses S3 to store code files in a robust and scalable manner
- [X] CRUD Fuctionalities for repos: Users can create, read, update and delete repos
- [X] Scalable Database: Snippet uses DynamoDB with a single table design for optimized reterival and storage of data

Additional possible features:
- [ ] Rate limiting using Redis or AWS Elastic cache
- [ ] Comments on repos like with Gist
- [ ] Ability for users to reset the password

<a id="architucture"></a>
## Architucture

**Technologies Used:**
- Go and Mux Router (Backend)
- DynamoDB(Database)
- AWS S3 (Storage)
- TailwindCSS (Styling)
- NextJS (Frontend)

<a id="license"></a>
## License

Licensed under [MIT License](./LICENSE)
