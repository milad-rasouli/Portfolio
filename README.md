# Portfolio

Portfolio aims you to have your personal website. It contains the following parts:

1. Home page where you can introduce yourself briefly.
2. Blog page where you can write and show them to people.
3. About Me page where you can write about yourself in details.
4. Contact page where others can write direct message for you as an admin.

## How To Run It

1. Docker

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

   - Run the Image

   ```bash
   sudo docker run -v "$(pwd)"/config.toml:/app/config.toml -p 80:5001 ghcr.io/milad75rasouli/portfolio:latest
   ```

   - Go to http://localhost:80 to visit your website

2. On your Machine. Follow these steps to run:

- Clone the project

```bash
git clone https://github.com/milad75rasouli/portfolio
```

- Get the packages:

```bash
 cd portfolio
 go mod tidy
```

- Install [templ](https://github.com/a-h/templ) and set it into the _PATH_.

```bash

```

- Install [just](https://github.com/just) and set it into the _PATH_.

```bash

```

- Run it

```bash
just run
```

## Contribution

I'd be happy if you contribute to this project. Please feel free to open issues and send pull requests.
