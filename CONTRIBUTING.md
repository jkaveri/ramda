# Contributing to Ramda

Thank you for your interest in contributing to Ramda! This document provides guidelines and steps for contributing to the project.

## Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct. Please be respectful and constructive in your interactions with other contributors.

## How to Contribute

1. Fork the repository
2. Create a new branch for your feature or bug fix
3. Make your changes
4. Add tests for new functionality
5. Update documentation if needed
6. Submit a pull request

## Development Setup

1. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/ramda.git
   cd ramda
   ```

2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bug-fix
   ```

3. Make your changes and commit them:
   ```bash
   git add .
   git commit -m "Description of your changes"
   ```

4. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

5. Create a Pull Request from your fork to the main repository

## Code Style Guidelines

- Follow Go standard formatting (`go fmt`)
- Write clear and descriptive commit messages
- Add comments for complex logic
- Keep functions pure and side-effect free
- Use meaningful variable and function names
- Write tests for new functionality

## Testing

- Add tests for any new functionality
- Ensure all tests pass before submitting a PR
- Run tests locally:
  ```bash
  go test ./...
  ```

## Documentation

- Update README.md if adding new functions
- Add examples for new functionality
- Keep documentation clear and concise
- Include type information in function documentation

## Pull Request Process

1. Update the README.md with details of changes if needed
2. Update the documentation with any new functions
3. The PR will be merged once you have the sign-off of at least one maintainer

## Questions?

If you have any questions, feel free to open an issue or contact the maintainers.

Thank you for contributing to Ramda!