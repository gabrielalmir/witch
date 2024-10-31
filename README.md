# Witch

Witch is a command-line application that allows you to search for trending repositories on GitHub. It is built using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework for creating rich terminal user interfaces.

## Features

- Search for GitHub repositories by keyword.
- Navigate through search results using arrow keys.
- View repository details such as name, description, stars, and owner.

## Installation

To install Witch, you need to have [Go](https://golang.org/dl/) installed on your machine. Then, you can clone the repository and build the application:

```sh
git clone https://github.com/gabrielalmir/witch.git
cd witch
go build
```

## Usage

Run the application from the command line:

```sh
./witch
```

Once the application is running, you can enter a keyword to search for repositories and navigate through the results using the arrow keys.

## Key Bindings

- `Enter`: Search for repositories.
- `Ctrl+C`, `q`, `Esc`: Quit the application.
- `Backspace`: Delete the last character of the search query.
- `Up`, `Down`: Navigate through the search results.
- `Left`, `Right`: Navigate through the pages of search results.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Any new ideas?

Contributions are welcome! Please open an issue or submit a pull request if you would like to contribute to this project.

## Author

Gabriel Almir - [github.com/gabrielalmir](https://github.com/gabrielalmir)
