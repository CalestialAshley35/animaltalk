# 🐾 Contributing to AnimalTalk

We moo-tastically welcome contributions from developers of all experience levels! Before participating, please read our [Code of Conduct](CODE_OF_CONDUCT.md).

## Table of Contents
1. [Getting Started](#🚀-getting-started)
2. [Development Workflow](#🛠️-development-workflow)
3. [Feature Requests](#✨-feature-requests)
4. [Bug Reports](#🐛-bug-reports)
5. [Code Contributions](#💻-code-contributions)
6. [Testing](#🧪-testing)
7. [Documentation](#📖-documentation)

## 🚀 Getting Started

1. **Fork** the repository
2. **Clone** your fork:
   ```bash
   git clone https://github.com/your-username/animaltalk.git
   ```
3. **Setup Development Environment**:
   ```bash
   cd animaltalk
   go mod tidy
   ```

## 🛠️ Development Workflow

### Branch Strategy
- `main`: Stable releases
- `dev`: Active development branch
- Feature branches: `feat/description` (e.g., `feat/add-penguin`)
- Fix branches: `fix/issue-number` (e.g., `fix/tts-linux-45`)

### Commit Message Convention
```
<type>(scope): <description>

[optional body]

[optional footer]
```
- **Types**: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`
- **Scope**: `cow`, `dog`, `tts`, `ascii`, `regex`, etc.

## ✨ Feature Requests

1. Check existing [issues](https://github.com/CalestialAshley35/animaltalk/issues)
2. Create new issue using "Feature Request" template
3. Include:
   - Use case description
   - Proposed implementation strategy
   - Compatibility with existing animals
   - Suggested ASCII art examples (for new animals)

## 🐛 Bug Reports

Use the bug report template and include:
- Animal type and conversation context
- Exact user input that triggered the issue
- Expected vs actual behavior
- OS and Go version (`go version` output)
- Stack trace (if available)

**Example Report**:
```markdown
## Environment
- OS: Ubuntu 22.04 LTS
- Go: 1.21.4
- Animal: Cow with TTS enabled

## Steps to Reproduce
1. Select cow with TTS
2. Type "dance party"
3. Observe crash

## Expected Behavior
Cow performs ASCII dance animation

## Actual Behavior
Segmentation fault at line 147 of cow_tts.go
```

## 💻 Code Contributions

### Pull Request Process
1. Sync with `dev` branch
2. Run validation suite:
   ```bash
   go fmt ./...
   go vet ./...
   go test -v ./...
   ```
3. Create PR to `dev` branch with:
   - Description of changes
   - Screenshots/ASCII examples (if visual change)
   - Update to [RESPONSES.md](RESPONSES.md) if modifying conversations

### Code Style Requirements
- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Maximum function length: 50 lines
- Regex patterns must use `\b` word boundaries
- ASCII art alignment preserved at 80-column width
- TTS implementations must include Linux fallback

## 🧪 Testing

### Response Pattern Validation
1. Add test cases to `animal_test.go`
2. Verify pattern priority order:
   ```go
   // High priority patterns first
   {`(?i)\bdance\b`, cowDanceResponse},
   {`(?i)\bhello\b`, cowGreetingResponse},
   ```

### TTS Requirements
- Linux contributors: Test with and without `espeak`
- Non-Linux: Verify text fallback works

## 📖 Documentation

1. Update [RESPONSES.md](RESPONSES.md) for new conversation patterns:
   ```markdown
   ### New Animal Responses
   | Trigger Phrase | Example Response |
   |----------------|------------------|
   | "fly"          | "Flap flap!"     |
   ```
2. Keep ASCII art documentation consistent
3. Update installation instructions for new dependencies

---

🙌 **First Time Contributors** are welcome! Look for issues tagged `good first issue`.

By contributing, you agree to license your work under the [MIT License](LICENSE).