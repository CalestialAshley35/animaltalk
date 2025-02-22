# AnimalTalk v6.0.0 🦍

Welcome to **AnimalTalk**! This is a fun and interactive chatbot that lets you chat with various animals such as cows, dogs, cats, and fish. Each animal has its own unique personality, fun responses, and even an ASCII art representation. 

## Features

- **Animal Interactions:** Choose from a variety of animals, each with their own fun responses.
- **Text-to-Speech Option:** Have your animal speak their response aloud (requires `espeak`).
- **Regex-Based Responses:** The chatbot uses regular expressions to identify keywords and give fun, tailored responses based on your input.
- **Multiple Animals to Chat With:** Includes Cow, Dog, Cat, and Fish, each with their own set of rules and responses.
- **Interactive Console:** Continuously interact with your chosen animal until you decide to exit.

## Installation

To run **AnimalTalk**, follow these steps:

1. **Clone the repository**:

    ```bash
    git clone https://github.com/CalestialAshley35/animaltalk.git
    ```

2. **Navigate to the project directory**:

    ```bash
    cd animaltalk
    ```

3. **Run the Go program**:

    ```bash
    go run animaltalk.go
    ```

4. **Follow the prompts** to interact with your chosen animal.

## How to Use

1. **Choose Your Animal:**
   - After starting the program, you will be prompted to choose an animal. You can type `cow`, `dog`, `cat`, or `fish`.

2. **Animal's ASCII Art and Prompt:**
   - Once you choose an animal, you’ll see their ASCII art and their unique prompt (e.g., `cow>`, `dog>`, etc.).

3. **Chat with the Animal:**
   - Type your message and the animal will respond based on predefined rules. If the animal recognizes keywords in your input, it will respond accordingly.
   - If the input isn't recognized, the animal will ask you to try again.

4. **Text-to-Speech Option:**
   - If you choose the **cow**, you will be asked if you want to enable text-to-speech. If you answer `y`, the response will be spoken aloud using the `espeak` command.

5. **Exit:**
   - To exit the chatbot, simply type `exit`.

## Animal Responses

Each animal has a set of predefined responses based on regex patterns that match user input. For example:

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

## Full Responses

For a full list of responses for each animal, please refer to `responses.md`. This document contains all of the keywords and matching responses for each animal, as well as other useful details.

## Requirements

- **Go 1.x+**: This application is written in Go and requires a Go environment to run.
- **espeak (optional)**: If you want to enable text-to-speech for the cow, ensure `espeak` is installed on your machine.

### Installing `espeak` on Linux

```bash
sudo apt-get install espeak
```


### Installing espeak on macOS

```
brew install espeak
```

## Installing espeak on Windows

Follow the espeak installation guide of Windows 

## License

This project is licensed under the MIT License - see the )ICENSE file for details.

---

Enjoy chatting with your new animal friends and have fun with AnimalTalk! 🐄🐕🐱🐟
