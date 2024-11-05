# Junior DevOps Assignment Submission

## Issues Found and Fixed

1. Missing .env file
   - Created .env.example file with necessary configuration
   - Added PORT and REDIS_URL settings

2. Docker-compose
   - Added Redis service to docker-compose
   - Configured web service to depend on Redis
   - Set up proper Redis connection URL
   - Corrected host to container port mapping
   - Mounted in the .env

3. Dockerfile
   - Exposed port for readability
   - explicitly tell that it is main.go that is built. Also for readability

4. Security
   - Added users.json to .gitignore as to not have secrets in public

5. Test
   - Added a simple unit test for ValidUser() with custom users (since users.json is not on git)
   - Navigate to src/tests
   - Run `go test -v`

6. CI
   - Added github actions to the project. /.github/workflows/go.yml
   - Slightly modified the "Go" template to use the same go version as the container,
   and points to the main.go for the build.
   - Added specific path for caching dependencies, as gh actions always looks in root for go.sum and go.mod. For this reason I set the working dir to be src.
   - Added the singular unit test to also run in the CI process.
   - The GitHub Actions workflow  performs:
   1. Sets up Go 1.22.7
   2. Downloads dependencies
   3. Runs unit tests
   4. Builds the application

7. K8s
   - Added deployment configuration in /k8s/deployment.yml
   - Deploys the "chat-app" 
   - Configuires the two necesdsary variables: port and redis_url - albeit explicitly
   - Sets up container configuration

After this I ran out of time, as I have not yet tried Kubernetes so it took some time to just get the deployment.yml done.
While I have created the deployment file, I have not been able to test it, as the deployment also depends on a redis service and a Kubernetes service? The Kubernetes service I didn't quite have the time to understand. It seemed like it is responsible for matching pods with containers with certain labels which is necessary when the cluster is running.

## How to Run

1. Create the .env file in the root directory. Give values to the variables seen in .env.example. (In this case port = 8080 and redis url = redis://redis:6379)
2. Create a users.json file in src/secrets. Inside put a list of strings of the accepted usernames. e.g: ["admin", "username"]
3. Run the application:
   bash:
   docker compose up
4. http://localhost:8080

## Actions status
![Go](https://github.com/SvenningsenDev/junior-devops-assignment-fork/actions/workflows/go.yml/badge.svg)