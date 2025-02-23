# 🐾 AnimalTalk - Interactive Animal Chat Simulator (v6.0.0)

<div align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License">
  <img src="https://img.shields.io/badge/Platform-Linux%2FmacOS%2FWindows-lightgrey" alt="Platform">
</div>

A charming CLI-based chatbot that lets you converse with various animals using regular expression pattern matching and ASCII art. Each animal has its unique personality and response patterns!

## 🌟 Features

- **Multi-Animal Support**: Chat with 4 different animals:
  - 🐄 Philosophical Cow 
  - 🐕 Energetic Dog
  - 😼 Sassy Cat
  - 🐠 Chill Fish
- **Regex-Powered Responses**: Intelligent pattern matching for dynamic conversations
- **ASCII Art Personalities**: Unique visual representation for each animal
- **Text-to-Speech**: Optional TTS support for Cow conversations (Linux only)
- **Contextual Understanding**: 100+ predefined response patterns
- **Interactive CLI**: Simple and intuitive text-based interface

## 🚀 Installation Guide

### Prerequisites
To get started, make sure you have the following prerequisites installed:

- **Go** version 1.21+  
- **TTS Support (for Cow Conversations)**:
  - **Linux**: Install `espeak` using:  
    ```bash
    sudo apt-get install espeak
    ```
  - **Windows/macOS**: Follow the installation steps from the [eSpeak official website](http://espeak.sourceforge.net/).

---

### Installation Instructions by Platform

#### **Linux**  
1. Clone the repository:
   ```bash
   git clone https://github.com/CalestialAshley35/animaltalk.git
cd animaltalk
go run animaltalk.go
```

- macOS: Download the binary from the [official release](https://github.com/CalestialAshley35/animaltalk/releases/tag/v6.0.0-for-macos)
- Windows: Download the binary from the [official release](https://github.com/CalestialAshley35/animaltalk/releases/tag/v6.0.0)

## 🕹️ Usage

1. Start the program:
   ```bash
   go run animaltalk.go
   ```
2. Choose an animal from the list
3. Begin your conversation!
4. Type `exit` at any time to quit

**Example Session:**
```text
Choose your animal (cow, dog, cat, fish): cow
Do you want Text to Speech? (y/n): n

        (__)
         (oo)
  /-------\/
 / |     ||
*  ||----||
   ^^    ^^
cow> hello
Moo! Moo! I'm Mr. Cow 🐄🐮, the barn's life of the party!
cow> 
```

## 📜 Responses List

Here are few responses:
- **Cow Responses:**  
  - "Hello" → "Moo! Moo! I'm Mr. Cow 🐄🐮, the barn's life of the party!"
  - "Dance" → "You bet! Watch me swing those hips in a moo-tastic two-step! 🕺💃"

- **Dog Responses:**  
  - "Hello" → "Woof! Woof! I’m the top dog! Bow-wow! 🐕"
  - "Bone" → "Bone! Bone! I could bury bones all day! 🦴"

- **Cat Responses:**  
  - "Hello" → "Meow! I’m the purr-fect companion. And, no, I don’t do tricks. 😼"
  - "Fish" → "Fish? YES! I WILL EAT ALL THE FISH! 🐟"

- **Fish Responses:**  
  - "Hello" → "Blub! Blub! I'm just swimming by, no biggie. 🐠"
  - "Water" → "Water? It’s my entire world, you know! 🌊"
 
All animal responses and conversation patterns are documented in [RESPONSES.md](RESPONSES.md). This includes:
- All possible animal responses

## 🔊 Text-to-Speech Feature

Exclusive to Cow conversations:
- Enabled with `y` during cow selection
- Uses `espeak` for Linux systems
- Adds vocal dimension to interactions
- Perfect for accessibility purposes
- if you type n Cow will response in text rather than voice 

## 🛠️ Technical Details

### Architecture
```mermaid
graph TD
    A[User Input] --> B{Animal Selection}
    B -->|Cow| C[TTS Check]
    B -->|Other Animals| D[Direct Interaction]
    C --> E[ASCII Display]
    E --> F[Input Processing]
    F --> G[Regex Matching]
    G --> H[Response Generation]
    H --> I[Output/TTS]
```

### Pattern Matching System
- Case-insensitive regex patterns
- Word boundary detection (`\b` anchors)
- Priority-based first-match system
- Fallback to default response

## 🤝 Contributing

We welcome contributions! Please see our [CONTRIBUTING.md](CONTRIBUTING.md) for:
- Feature request guidelines
- Bug report templates
- Code style requirements
- Pull request processes

## 📄 License

MIT License - See [LICENSE](LICENSE) for details.

## 🐮 Acknowledgments

- ASCII art from various online sources
- Emoji support from Unicode Consortium
- TTS powered by eSpeak NG
```