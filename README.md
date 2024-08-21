# Remove Node Modules Script

This Go script recursively deletes all `node_modules` directories in the current directory and its subdirectories.

## Usage

1. **Clone or download the script** to your local machine.

2. **Navigate to the directory** where the script is located using the terminal.

3. **Run the script** with the following command:

    ```bash
    go run remove-node-modules.go
    ```

   This command will remove all `node_modules` directories found in the directory where the script is executed, including all subdirectories.

## How It Works

- The script traverses the directory tree starting from the directory where it is executed.
- For every subdirectory it finds named `node_modules`, it deletes the folder and all its contents using the `os.RemoveAll` function.
- The script skips further traversal into any `node_modules` directory after deleting it, ensuring that the removal is efficient.

## Prerequisites

- **Go Programming Language**: Ensure that Go is installed on your machine. You can download it from the official [Go website](https://golang.org/).

## Warning

This script permanently deletes all `node_modules` directories and their contents within the directory from which the script is executed. Use with caution.

## License

This script is provided as-is, without any warranty. Feel free to modify and use it according to your needs.
