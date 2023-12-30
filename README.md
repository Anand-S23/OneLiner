# One Liner

Twitter like collerbative social media platform for writers

## Table of Contents 

- [Quick Start](#quick-start)
- [Features](#features)
- [Architucture](#architucture)
- [License](#license)

<a id="quick-start"></a>
## Quick Start

One Liner uses docker for easy development, docker and docker-compose are pre-requistes. Follow the following steps once the repository is cloned locally:

1. Create a `.env` file using example.env as an base
2. Run `docker-compose up --build`

Note: In order to test out if this is working as expected you can `curl http://localhost:8080/ping`, where the result should be `pong`

<a id="features"></a>
## Features

The core idea for the platform, is to allow users to collaberate on short stories. Users can either write a prompt for a short story, or write a story on a given prompt.

- [ ] Authentication System - allow users to register/login/logout
- [ ] Prompts - allow users to do crud operations on prompts
- [ ] Stories - allow users to do crud operations on stories
- [ ] Profile - allow users to view posts on one person
- [ ] Follow System - allow users to follow other users to get prompts/stories in their feed

<a id="architucture"></a>
## Architucture

**Technologies Used or Will Use:**
- Go and Mux Router (Backend)
- DynamoDB(Database)
- ElastiCache (Rate Limiting + Cache)
- Ozzo (Input validation)
- AWS S3 (Storage)
- TailwindCSS (Styling)

<a id="license"></a>
## License

Licensed under [MIT License](./LICENSE)
