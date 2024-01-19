# Sai - Shell AI Assistant

Sai is a command-line tool designed to enhance your shell experience by providing intelligent and context-aware command suggestions using OpenAI's GPT-4 model. Whether you're working with git, navigating directories, or performing complex shell operations, Sai offers a helping hand.

## Features

- **AI-Powered Command Suggestions:** Leverages the power of GPT-4 to understand and suggest relevant shell commands.
- **Context-Aware:** Takes into account your current directory, operating system, and recent command history to provide accurate suggestions.
- **User-Friendly:** Simple and intuitive to use, integrating seamlessly into your existing shell workflow.

## Installation

To install Sai, follow these steps:

```bash
# Clone the repository
git clone https://github.com/lguibr/sai.git

# Change directory
cd sai

# Run the installation script
./install.sh
```

Ensure that Go is installed on your system before running the installation script.

## Usage

![Example Usage](./examples.gif)

To use Sai, simply type sai followed by your query in quotes. For example:

```
sai "how to undo the last git commit"
git reset HEAD~1
```

## Configuration

Before using Sai, you need to set your OpenAI API key as an environment variable:

```bash

export OPENAI_API_KEY='your-api-key-here'
```

You may add this line to your .bashrc, .zshrc, or equivalent shell configuration file to make it persistent.

Contributions
Contributions are welcome! If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

Licensing
The code in this project is licensed under MIT license. See LICENSE for more information.
