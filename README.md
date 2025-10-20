# Go File Organizer CLI

A simple and efficient command-line tool built with Go to automatically organize files in a directory into subdirectories based on their file extensions.

---

## 🚀 Features

- **Automatic Organization**: Scans a target directory and moves files into folders named after their extensions (e.g., `.jpg` files go into a `jpg` folder).
- **Intelligent Handling**: Correctly ignores any subdirectories and files that do not have an extension.
- **Safe Directory Creation**: Automatically creates destination directories if they don't already exist.
- **Cross-Platform**: Built with Go's standard library to run on Windows, macOS, and Linux.
- **Simple to Use**: Takes a single command-line argument for the directory path.

---

## 📋 Prerequisites

- [Go](https://go.dev/doc/install) (Version 1.24 or later) installed on your system.

---

## ⚙️ How to Use

1.  Clone or download the project.
2.  Navigate to the project directory in your terminal.
3.  Run the program using `go run`, providing the path to the directory you want to organize as a command-line argument.

### Example Command

```bash
# To organize files in a directory located at /home/user/Downloads
go run main.go /home/user/Downloads

### ⚠️ **Safety Warning**

This tool permanently moves files on your filesystem. It is highly recommended to **first test it on a sample directory** with copied files before running it on important directories.

---

## ✨ Future Improvements

- **Dry Run Mode**: Add a `--dry-run` flag to show which files would be moved without actually performing the operation.
- **Custom Grouping**: Allow users to define custom rules for grouping extensions (e.g., move `.jpg`, `.png`, and `.gif` into a single `images` folder).
- **Unit Tests**: Add table-driven tests for the core organization logic.