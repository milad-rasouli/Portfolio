# Portfolio

Portfolio aims for you to have your personal weblog. It contains the following parts:

1. Home page where you can introduce yourself briefly.
2. Blog page where you can write and show them to people.
3. About Me page where you can write about yourself in detail.
4. Contact page where others can write direct messages for you as an admin.

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

   - Copy the configs into _config.toml_

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
> if you clone the project without its submodules please run this you get them ```git submodule update --init --recursive ```
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

Then go to the shown address in your browser.

## Contribution

I'd be happy if you contribute to this project. Please feel free to open issues and send pull requests.
