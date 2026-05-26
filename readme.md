# Gotermi 🐹🌸

> A cute terminal pet that lives inside your command line.

Gotermi is a Go-based CLI pet app where users can create a virtual pet, interact with it, track stats, gain XP, level up, and save progress locally.

Built for people who want their terminal to feel a little less boring and a little more alive.

---

## Preview

```bash
Welcome to Gotermi 🐹
Your terminal pet is alive.

        /\_/\
       ( o.o )
        > ^ <

   ─────────────
      CHUNARI🌸
   🌸🌸🌸
   ─────────────

╭────────────────────╮
  Chunari 🌸
├────────────────────┤
  Level      1
  XP         ●○○○○ 13%
  Hunger     ●●○○○ 55%
  Happy      ●●●○○ 75%
  Energy     ●●●●● 100%
├────────────────────┤
  Keep alive. No excuses 🐹
╰────────────────────╯
```

---

## What is Gotermi?

Gotermi is a small terminal pet game built with Go.

Users can create their own pet, feed it, play with it, let it sleep, check its mood, and continue progress anytime because the pet data is saved locally.

It runs directly inside the terminal and can also be installed as a global command, so users can open Gotermi from anywhere.

---

## Features

| Feature | Description |
|---|---|
| Terminal Pet | A virtual pet that runs inside the command line |
| Custom Pet Name | Users can name their own pet on first run |
| ASCII Pet UI | Cute terminal-based pet visuals |
| Pet Stats | Tracks hunger, happiness, energy, XP, and level |
| Mood System | Pet mood changes based on its stats |
| XP System | Users gain XP by interacting with the pet |
| Level System | Pet levels up as XP increases |
| Random Events | Small random changes make the pet feel alive |
| JSON Save System | Pet progress is saved locally |
| Auto Save | Data saves automatically after actions |
| Safe Exit | Handles exit without losing progress |
| Global Command | Can be installed and used as `gotermi` from anywhere |

---

## Tech Stack

| Area | Technology |
|---|---|
| Language | Go |
| Storage | JSON |
| Interface | CLI |
| UI | ASCII Art |
| Terminal Styling | ANSI Colors |
| Data Handling | Go Structs + File Handling |

---

## Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
cd YOUR_REPO_NAME
```

Example:

```bash
git clone https://github.com/amritkang165/Gotermi.git
cd Gotermi
```

### 2. Run Gotermi

```bash
go run main.go
```

On the first run, Gotermi will ask for a pet name.

```bash
Name your pet:
Chunari
```

After that, your pet will be created and saved automatically.

---

## Install Gotermi as a Terminal Command

To use Gotermi from anywhere in your terminal, build and install it:

```bash
go build -o gotermi
sudo mv gotermi /usr/local/bin/
```

Now run:

```bash
gotermi
```

Example:

```bash
cd ~
gotermi
```

After editing the code, rebuild and reinstall:

```bash
go build -o gotermi
sudo mv gotermi /usr/local/bin/
```

---

## Usage

After starting Gotermi, users can choose actions from the terminal menu.

Example actions:

```bash
1. Feed pet
2. Play with pet
3. Let pet sleep
4. Check status
5. Exit
```

Each action changes the pet stats.

| Action | Effect |
|---|---|
| Feed | Reduces hunger |
| Play | Increases happiness and uses energy |
| Sleep | Restores energy |
| Status | Shows current pet stats |
| Exit | Saves and closes the app |

---

## Saved Data

Gotermi saves pet data locally at:

```bash
~/.gotermi/pet.json
```

Example saved data:

```json
{
  "name": "Chunari",
  "hunger": 55,
  "happiness": 80,
  "energy": 95,
  "level": 1,
  "xp": 28
}
```

---

## Reset Pet

To reset the current pet and create a new one:

```bash
rm ~/.gotermi/pet.json
```

Then run Gotermi again:

```bash
gotermi
```

or:

```bash
go run main.go
```

---

## Common Commands

| Task | Command |
|---|---|
| Run from source | `go run main.go` |
| Build executable | `go build -o gotermi` |
| Run built executable | `./gotermi` |
| Install globally | `sudo mv gotermi /usr/local/bin/` |
| Run globally | `gotermi` |
| View saved data | `cat ~/.gotermi/pet.json` |
| Reset pet | `rm ~/.gotermi/pet.json` |
| Check Go version | `go version` |

---

## Project Structure

```bash
Gotermi/
├── main.go
└── README.md
```

After running Gotermi:

```bash
~/.gotermi/
└── pet.json
```

---

## How It Works

Gotermi stores pet information inside a Go struct.

```go
type Pet struct {
    Name      string `json:"name"`
    Hunger    int    `json:"hunger"`
    Happiness int    `json:"happiness"`
    Energy    int    `json:"energy"`
    Level     int    `json:"level"`
    XP        int    `json:"xp"`
}
```

When the app starts, it loads saved pet data from the JSON file.

When users perform actions, Gotermi updates the pet stats and saves the new data automatically.

---

## Learning Goals

This project demonstrates:

- Building a CLI app in Go
- Working with structs
- Using pointers
- Handling terminal input
- Creating a menu-based program
- Saving and loading JSON data
- Managing local files and folders
- Using ANSI colors in terminal output
- Building a Go executable
- Installing a custom terminal command

---

## Roadmap

Possible future improvements:

- More pet types
- Multiple pet support
- More random events
- Daily streak system
- Pet health system
- More moods and expressions
- Custom themes
- Better animations
- Sound effects
- Mini games inside the terminal

---

## Contributing

Contributions are welcome.

To contribute:

1. Fork the repository
2. Create a new branch
3. Make your changes
4. Commit your work
5. Push your branch
6. Open a pull request

```bash
git checkout -b feature/your-feature-name
git add .
git commit -m "Add your feature"
git push origin feature/your-feature-name
```

---

## License

This project is open source.

You can add a license file such as MIT License if you want to define usage and contribution rules clearly.

---

## Final Result

After setup, users can run:

```bash
gotermi
```

from anywhere in their terminal and take care of their own tiny terminal pet.

Built with Go, JSON, ASCII art, and terminal chaos. 🐹🌸