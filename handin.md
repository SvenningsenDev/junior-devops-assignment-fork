# Junior DevOps Assignment Submission

## Issues Found and Fixed

1. Missing .env file
   - Created .env file with necessary configuration
   - Added PORT and REDIS_URL settings

2. Missing Redis dependency
   - Added Redis service to docker-compose
   - Configured web service to depend on Redis
   - Set up proper Redis connection URL

## How to Run

1. Make sure the .env file is in the root directory
2. Create a users.json file in src/secrets. Inside put a list of strings of the accepted usernames. e.g: ["admin", "username"]
2. Run the application:
   ```bash
   docker compose up