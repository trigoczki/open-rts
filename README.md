# Modular RTS Game Server

An open-source, modular real-time strategy (RTS) game server built in Go. This project is designed to be fully data-driven, allowing different game mechanics to be defined using structured rulesets in JSON format. The server is built incrementally, starting with basic city-building mechanics, and expanding to include colonization, trade, combat, and diplomacy.

## Features

- 🚀 **Modular Ruleset** – *(Planned)* - Define game mechanics using a structured configuration file.
- 🏙 **City Building** – *(Planned)* - Establish cities and manage resources.
- 🌍 **Map-Based Gameplay** – *(Planned)* - Supports a tile-based world.
- 🤝 **Trade & Economy** *(Planned)* – Exchange resources between players.
- ⚔ **Combat & Warfare** *(Planned)* – Engage in battles and strategic conquests.
- 🔄 **Incremental Development** – Built in stages to ensure flexibility.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (1.23+ recommended)
- Git

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/trigoczki/open-rts
   cd open-rts
   ```

2. Install dependencies:

   ```sh
   go mod tidy  
   ```

3. Run the server:

   ```sh
   go run open_rts.go
   ```