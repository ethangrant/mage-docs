# mage-docs

mage-docs is a CLI tool written in Go, designed to generate a `README.md` file for Magento 2 modules. By extracting key information directly from the module, mage-docs helps developers quickly create comprehensive and consistent documentation. It can also be useful for reviewing new modules providing a high level overview.

## Features

mage-docs can document the following aspects of a Magento 2 module:

- Module name and its dependencies
- Routes and API routes
- Utilised layouts
- Implemented mixins
- Database schemas
- Events dispatched
- Registered observers
- Plugins
- Preferences

## Installation

mage-docs is written in Go and has been compiled into an executable. See the latest releases to download mage-docs.

## Usage

```bash
./mage-docs
```

When you run mage-docs, you will be guided through a series of prompts to configure the documentation process:

1. **Enter the absolute path to the Magento 2 module:**
2. **Enter the name of the output file:**
3. **Select what you want to document:**

**Command TUI:**
![mage-docs-tui](readme/images/mage-docs-tui.png)

**Generated .md file:**
![mage-docs-generated](readme/images/mage-docs-generated-readme.png)

## Sample README.md

See sample README generated by mage-docs [README_SAMPLE.md](README_SAMPLE.md)

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or improvements, create an issue.

## License

mage-docs is open-sourced software licensed under the [MIT license](LICENSE).