# Contributing to ASCII-Art

Thank you for considering contributing to this project!

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR-USERNAME/ascii-art.git
   cd ascii-art
   ```
3. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```

## Development Setup

**Requirements:**
- Go 1.24 or higher

**Setup:**
```bash
go mod tidy
```

## Making Changes

1. **Follow Go conventions:**
   - Format code: `go fmt ./...`
   - Keep code simple and readable

2. **Add tests for new features:**
   - Unit tests in `internal/ascii/`
   - Edge case tests in `edge_cases_test.go`

3. **Update documentation:**
   - Update README.md if needed
   - Add entry to CHANGELOG.md

4. **Ensure tests pass:**
   ```bash
   go test ./...
   ```

## Testing

Run tests with:
```bash
make test           # Run all tests
go test ./...       # Alternative
```

Test the binary:
```bash
make build
./ascii-art "Test"
```

## File Structure

- `cmd/ascii-art/` - Main application entry point
- `internal/ascii/` - Core logic (art generation, filtering, etc.)
- `internal/assets/` - Banner font files

## Submission Process

1. **Test your changes:**
   ```bash
   go test ./...
   go build ./cmd/ascii-art
   ```

2. **Format your code:**
   ```bash
   go fmt ./...
   ```

3. **Commit with clear messages:**
   ```bash
   git commit -m "feat: add new feature"
   ```
   Use conventional commit prefixes: `feat:`, `fix:`, `docs:`, `test:`, `refactor:`, `chore:`

4. **Push to your fork:**
   ```bash
   git push origin feature-name
   ```

5. **Create a Pull Request:**
   - Provide a clear description
   - Reference any related issues
   - Ensure all tests pass

## Code Review

- All submissions require review
- Tests must pass
- Code must follow Go conventions
- Documentation must be updated

## Questions?

Open an issue for questions or suggestions!
