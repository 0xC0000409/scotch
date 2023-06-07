<a name="readme-top"></a>
# Scotch


## About The Project

This project aims to develop and combine various Golang libraries in order to create starter project for creating RESTful APIs in golang.


## Getting Started


### Prerequisites

To start a project with a docker you would require `docker`. In order to start project manually, `go >= 1.20` is needed.


### Installation with docker (recommended)

1. Clone the repo:
   ```sh
   git clone git@github.com:0xC0000409/scotch.git
   ```
2. Create `.env` file in the project root and copy the contents of `.env.example` into it.
3. Start the project:
   ```sh
   docker compose up
   ```


### Manual installation

1. Clone the repo:
   ```sh
   git clone git@github.com:0xC0000409/scotch.git
   ```
2. Create `.env` file in the project root and copy the contents of `.env.example` into it.
3. Install air:
   ```sh
   go install github.com/cosmtrek/air@latest
   ```
4. Install project dependencies:
   ```sh
   go get .
   ```
5. Start the project with air:
   ```sh
   air
   ```


## Roadmap

- [x] Add readme
- [x] Add license
- [ ] Add permission schema
- [ ] Add websocket routing
- [ ] Add i18n support


## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


## License

Distributed under the MIT License. See `LICENSE.txt` for more information.


## Acknowledgments

* [Go](https://go.dev/)
* [Gin](https://github.com/gin-gonic/gin)
* [Air](https://github.com/cosmtrek/air)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
