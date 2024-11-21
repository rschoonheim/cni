# Container Network Interface (CNI)

Container Network Interface (CNI) is an abstraction layer for network namespaces in Linux. It can be used by container
runtimes to manage network namespaces and network interfaces.

## Directory Structure

The project follows the following directory structure:

```bash
|-- src/ # Source code
  |-- ./ # Facade like functions that are meant to be used by the package users
  |-- internal/ # Modules that are not meant to be used outside the package
|-- test/ # Mirror of the src/ directory but for tests
```