
# Backend for blogs

This project has been made to provide a backend for blog websites. It is a modern and cloud native backend application. It is made with a CRUD API easy to use.




## Authors

- [@dzenancindrak](https://www.github.com/DzeCin)


## Roadmap

- Add authentication with OAuth2/OIDC

- Add comments

- Add likes

- Full test coverage


## Environment Variables

To run this project, you will need some env variables.

If you want to develop, add variables in a .development.env file in the root folder.

`DB_USERNAME` username to be used with the mongoDB database.

`DB_PASSWORD` password for the db.

`DB_HOST` host where the db is hosted

`DB_NAME` the database name to use for the project.

If you want to run it in production, pass those env variable without .env file.
## Run Locally

Clone the project

```bash
  git clone https://github.com/DzeCin/blog-backend.git
```

Go to the project directory

```bash
  cd blog-backend
```

Build the docker image

```bash
  docker build -t blog:v1 .
```

Create a .development.env file and setup env variables

Start the docker container

```bash
  docker run --env-file=.development.env --rm --publish 8080:8080 blog:v1
```


## Running Tests

To run tests, run the following command

```bash
  go test ./tests
```