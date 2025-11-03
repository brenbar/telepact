# Telepact VS Code Extension

This extension provides JSON schema validation for Telepact API schema files (`*.telepact.json`).

## Features

- **JSON Schema Validation**: Automatically validates `*.telepact.json` files against the Telepact JSON schema
- **IntelliSense**: Get autocomplete suggestions and inline documentation while editing schema files
- **Error Detection**: Catch syntax errors and schema violations as you type

## Installation

### From VSIX File

1. Download the latest `.vsix` file from the releases or build it from source
2. Open VS Code
3. Go to Extensions view (`Ctrl+Shift+X` or `Cmd+Shift+X`)
4. Click the `...` menu at the top of the Extensions view
5. Select "Install from VSIX..."
6. Choose the downloaded `telepact-*.vsix` file

### Building from Source

```bash
cd sdk/vscode
make
```

This will create a `telepact-*.vsix` file that you can install using the steps above.

## Usage

Once installed, the extension automatically activates when you open any file with the `.telepact.json` extension. The editor will:

- Show error underlines for invalid syntax
- Provide autocomplete suggestions for valid schema elements
- Display documentation on hover

### Example

Create a file named `example.telepact.json`:

```json
[
    {
        "///": "A simple struct definition",
        "struct.User_": {
            "name": "string",
            "age": "integer"
        }
    }
]
```

The extension will validate this file against the Telepact schema and provide IntelliSense support.

## About Telepact

Telepact is a multi-language API ecosystem built around a unified schema definition. Define your API once in `.telepact.json` files and use language-specific libraries to implement clients and servers.

Learn more at [github.com/telepact/telepact](https://github.com/telepact/telepact)

## License

Apache License 2.0 - See LICENSE file for details

