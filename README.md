# Go Bootcamp RecipeDB Analyzer

## Overview

Go Bootcamp RecipeDB Analyzer is a command-line application for analyzing recipe databases stored in XML and JSON formats. This tool reads, compares, and assesses changes in different recipe databases, highlighting modifications such as added, removed, or altered recipes and ingredients.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Getting Started](#getting-started)
4. [Usage](#usage)
    1. [Exercise 00: Reading](#exercise-00-reading)
    2. [Exercise 01: Assessing Damage](#exercise-01-assessing-damage)
    3. [Exercise 02: Afterparty](#exercise-02-afterparty)
5. [Project Structure](#project-structure)

## Introduction

There are many popular data formats in the world of programming, and Go makes it easy to work with them, particularly XML and JSON. This project simulates a bakery's recipe database stored in XML and compares it with a "stolen" database stored in JSON format. The application can detect various differences between the databases and highlight them.

## Features

- Read recipe databases in both XML and JSON formats.
- Compare two databases and detect changes, such as:
    - Added or removed cakes.
    - Changes in cooking time.
    - Added or removed ingredients.
    - Changes in ingredient quantities or units.
- Command-line interface for ease of use.

## Getting Started

### Prerequisites

- Go programming language installed (version 1.16 or higher recommended)
- Git for version control

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/go-bootcamp-recipeDB-analyzer.git
    ```
2. Navigate to the project directory:
    ```bash
    cd go-bootcamp-recipeDB-analyzer
    ```
3. Build the application:
    ```bash
    go build -o compareDB ./src/cmd/compareDB/main.go
    go build -o readDB ./src/cmd/readDB/main.go
    ```

## Usage

### Exercise 00: Reading

To read a database file and convert its format (from JSON to XML or vice versa), use the `readDB` command:

```bash
./readDB -f original_database.xml
./readDB -f stolen_database.json
```

This will output the contents of the specified file in the opposite format (JSON will be converted to XML, and vice versa).

### Exercise 01: Assessing Damage

To compare the original database with the stolen one and identify changes, use the compareDB command:

```bash
./compareDB --old original_database.xml --new stolen_database.json
```

Expected output format:
```rust
ДОБАВЛЕН торт "Маффин Лунный свет"
УДАЛЕН торт "Черничный кекс"
ИЗМЕНИЛОСЬ время готовки для торта "Торт Красный Бархат с клубникой" - "45 мин" вместо "40 мин"
ДОБАВЛЕН ингредиент "Кофейные зерна" для торта "Торт Красный Бархат с клубникой"
УДАЛЕН ингредиент "Экстракт ванили" для торта "Торт Красный Бархат с клубникой"
ИЗМЕНИЛАСЬ единица измерения для ингредиента "Мука" для торта "Торт Красный Бархат с клубникой" - "кружки" вместо "чашки"
ИЗМЕНИЛОСЬ количество для ингредиента "Клубника" для торта "Торт Красный Бархат с клубникой" - "8" вместо "7"
УДАЛЕНА единица измерения "шт" для ингредиента "Корица" для торта "Торт Красный Бархат с клубником"
```

### Exercise 02: Afterparty

To compare two filesystem dumps, use the compareFS command:

```bash
./compareFS --old snapshot1.txt --new snapshot2.txt
```

This will output any added or removed files between the two snapshots.

## Project Structure

```graphql
go-bootcamp-recipeDB-analyzer/
│
├── src/
│   ├── cmd/
│   │   ├── compareDB/       # Main command for comparing databases
│   │   └── readDB/          # Main command for reading databases
│   ├── data/                # Sample XML and JSON data files
│   ├── internal/
│   │   ├── converter/       # Code for data conversion (XML to JSON and vice versa)
│   │   └── dbreader/        # Database readers for XML and JSON
│   └── pkg/
│       └── utils/           # Utility functions (e.g., file handling, error handling)
│
└── .gitignore               # Ignore files for version control
```

### Explanation of the Directories

- **`cmd/`**: This directory contains the main application entry points. It has subdirectories for each command (e.g., `compareDB` and `readDB`), each with its own `main.go` file.

- **`data/`**: Stores sample databases in different formats (XML and JSON) for testing and development.

- **`internal/`**: This directory contains the core internal packages of the project, organized into:
    - **`converter/`**: Code for converting data between formats (e.g., JSON to XML).
    - **`dbreader/`**: Implements the `DBReader` interface for handling different file formats. This includes:
        - `dbreader.go`: Defines the interface for reading the databases.
        - `json_reader.go` and `xml_reader.go`: Implementations for reading JSON and XML databases.
        - `types.go`: Defines common data structures.

- **`pkg/`**: Contains general-purpose utilities and helper functions.
    - **`utils/`**: Includes utility functions for error handling, file operations, and command-line parsing.

- **`testfiles/`**: A folder intended for storing test files, useful for testing the application.

- **`.gitignore`**: Specifies files and directories that Git should ignore (e.g., `.idea/`).

- **`go.mod`**: The Go module file, which contains information about the module and its dependencies.
