# Portfolio

Portfolio aims for you to have your personal weblog. It contains the following parts:

1. **Home Page:** A space for you to introduce yourself briefly.
2. **Blog Page:** A platform where you can write and showcase your thought to people.
3. **About Me Page:** A detailed section where you can write more about yourself.
4. **Contact Page:** A communication channel where others can write direct messages to you as an admin.

[Here is where you can see my plans for the project.](https://github.com/users/Milad75Rasouli/projects/5) If you have any suggestions to enhance this project, please feel free to open an issue.

## How To Run It

1. **Docker**

   - Pull the image

   ```bash
   docker pull ghcr.io/milad75rasouli/portfolio:latest
   ```

   - Create _config.toml_

   ```bash
    touch config.toml
   ```

   - Copy the configurations into _config.toml_

   ```config.toml
    debug = true
    admin_email = "milad@gmail.com"

    [database]
    is_sqlite=true
    connection_timeout="2s"

    [cipher]
    paper="verySecurePaper"
    time=2
    memory=131072 #128\*1024
    Thread=5

    [jwt]
    refresh_secret_key="key1234567890" # it should be over 48 characters long to be secure
    access_secret_key="key123456" # it should be over 32 characters long to be secure
   ```

   - Change the config file (admin_email, paper must be changed)

   - Run the Image

   ```bash
   sudo docker run -v "$(pwd)"/config.toml:/app/config.toml -p 80:5001 ghcr.io/milad75rasouli/portfolio:latest
   ```

   - Go to http://localhost:80 to visit your website

2. **On your Machine.** Follow these steps to run:

- Clone the project

```bash
git clone --recursive https://github.com/milad75rasouli/portfolio
```

> [!NOTE]
> if you clone the project without its submodules please run this you get them `git submodule update --init --recursive `

- Get the packages:

```bash
 cd portfolio
 go mod tidy
```

- Download and install [templ](https://github.com/a-h/templ/releases) and set it into the _PATH_.

- Download and install [just](https://github.com/casey/just/releases) and set it into the _PATH_.

- Run it

```bash
just run
```

## K8S

1. Run minikube:

```bash
minikube start
```

> [!TIP]
> You might need to send the image to minikube:
> ` eval $(minikube docker-env)`

2. Apply the file:

```bash
kubectl apply -f deployment.yml
```

> [!TIP]
> See your services:
> ` kubectl get services`
> See your pods;
> `kubectl get pods`

3. Expose the service to see the website:

```bash
minikube service portfolio-service
```

Then, visit the displayed address in your browser.

## Contribution

Your contributions to this project are welcome. Please feel free to open issues and send pull requests.
