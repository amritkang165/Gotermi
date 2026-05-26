# Gotermi 🐹🌸

Gotermi is a cute terminal pet built with Go.

You can feed your pet, play with it, let it sleep, check its mood, gain XP, level up, and save progress using JSON.

It runs directly inside your terminal.

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

## Features

- Terminal-based pet game
- ASCII pet UI
- Colored terminal output
- Hunger, happiness, and energy stats
- XP and level system
- Random events
- Mood-based pet expressions
- JSON save/load system
- Auto-save after actions
- Ctrl+C safe exit support
- Saves pet data locally
- Can be installed as a global terminal command

---

## Requirements

You need Go installed on your system.

Check if Go is installed:

```bash
go version
```

If Go is installed, you will see something like:

```bash
go version go1.xx.x darwin/arm64
```

If Go is not installed, download and install it from:

```bash
https://go.dev/dl/
```

After installation, close and reopen your terminal, then check again:

```bash
go version
```

---

## Installation and Setup

### 1. Clone the repository

Replace the URL below with the actual Gotermi GitHub repo link.

```bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
```

Example:

```bash
git clone https://github.com/amritkang165/Gotermi.git
```

---

### 2. Go inside the project folder

```bash
cd YOUR_REPO_NAME
```

Example:

```bash
cd Gotermi
```

---

### 3. Run Gotermi directly

```bash
go run main.go
```

This starts Gotermi directly from the source code.

On the first run, Gotermi will ask you to name your pet:

```bash
Name your pet:
```

Type your pet name and press Enter.

Example:

```bash
Chunari
```

After that, your pet will be created and saved automatically.

---

## Build the App

If you want to create a runnable executable file, use:

```bash
go build -o gotermi
```

This creates a file named:

```bash
gotermi
```

Now run it using:

```bash
./gotermi
```

---

## Install Gotermi as a Global Terminal Command

After building the app, you can install it globally so you can run Gotermi from anywhere in your terminal.

First build the app:

```bash
go build -o gotermi
```

Then move it to `/usr/local/bin`:

```bash
sudo mv gotermi /usr/local/bin/
```

Now you can run Gotermi from anywhere:

```bash
gotermi
```

Example:

```bash
cd ~
gotermi
```

If Gotermi opens, installation worked.

---

## What Does `sudo` Mean?

`sudo` means:

```bash
superuser do
```

It runs a command with admin permission.

This is needed because `/usr/local/bin/` is a protected folder.

When you run:

```bash
sudo mv gotermi /usr/local/bin/
```

your system may ask for your computer password.

Type your password and press Enter.

Important:

When typing the password, nothing may appear on screen. No dots, no stars, nothing.

That is normal. Just type the password and press Enter.

Use `sudo` carefully. Only use it when you understand what the command does.

---

## Why Move Gotermi to `/usr/local/bin/`?

Your terminal looks inside certain folders to find commands.

One of those folders is:

```bash
/usr/local/bin/
```

When you move the `gotermi` executable there, your terminal can find it from anywhere.

That means instead of running:

```bash
./gotermi
```

inside the project folder, you can simply run:

```bash
gotermi
```

from any folder.

---

## Where Pet Data Is Saved

Gotermi saves your pet data here:

```bash
~/.gotermi/pet.json
```

This keeps the saved pet data in your home directory instead of creating random `pet.json` files everywhere.

To view your saved pet data:

```bash
cat ~/.gotermi/pet.json
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

## How to Create Your Own Pet

Run Gotermi:

```bash
gotermi
```

Or, if you are running from source:

```bash
go run main.go
```

On first run, it will ask:

```bash
Name your pet:
```

Type your pet name.

Example:

```bash
Momo
```

Press Enter.

Gotermi will create your pet and save it automatically.

---

## How to Reset Your Pet

If you want to delete your current pet and create a new one, delete the saved JSON file:

```bash
rm ~/.gotermi/pet.json
```

Then run Gotermi again:

```bash
gotermi
```

Or:

```bash
go run main.go
```

Gotermi will ask you to create a new pet again.

---

## How to Update Gotermi After Changing Code

If you edit `main.go`, the global `gotermi` command will not update automatically.

Why?

Because the global command is a compiled binary stored here:

```bash
/usr/local/bin/gotermi
```

Your source code is here:

```bash
main.go
```

The flow is:

```bash
main.go → go build → gotermi executable → /usr/local/bin/gotermi
```

So after editing the code, rebuild and reinstall:

```bash
go build -o gotermi
sudo mv gotermi /usr/local/bin/
```

Then run:

```bash
gotermi
```

Now the latest version will run.

---

## Development Workflow

Use this while coding and testing:

```bash
go run main.go
```

Use this when you want to build the app:

```bash
go build -o gotermi
```

Use this when you want to install or update the global command:

```bash
sudo mv gotermi /usr/local/bin/
```

Complete workflow:

```bash
cd Gotermi
go run main.go
go build -o gotermi
sudo mv gotermi /usr/local/bin/
gotermi
```

---

## Common Commands

Run directly from source:

```bash
go run main.go
```

Build executable:

```bash
go build -o gotermi
```

Run built executable locally:

```bash
./gotermi
```

Install globally:

```bash
sudo mv gotermi /usr/local/bin/
```

Run globally:

```bash
gotermi
```

View saved pet data:

```bash
cat ~/.gotermi/pet.json
```

Reset pet:

```bash
rm ~/.gotermi/pet.json
```

Check Go version:

```bash
go version
```

---

## Project Structure

```bash
Gotermi/
├── main.go
└── README.md
```

After global installation, the executable is stored here:

```bash
/usr/local/bin/gotermi
```

Pet data is stored here:

```bash
~/.gotermi/pet.json
```

---

## How Gotermi Works

Gotermi uses a simple Go struct to store pet data:

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

The pet has:

- Name
- Hunger
- Happiness
- Energy
- Level
- XP

The data is saved into a JSON file so your pet does not reset every time you close the terminal.

---

## Why Are Struct Fields Capitalized?

In Go, JSON encoding works properly with exported fields.

Exported fields start with capital letters.

Correct:

```go
Name
Hunger
Happiness
Energy
```

Wrong:

```go
name
hunger
happiness
energy
```

If the fields are lowercase, Go’s JSON package cannot access them properly.

---

## Terminal Colors

Gotermi uses ANSI escape codes for colored terminal output.

Example:

```go
const Green = "\033[32m"
const Reset = "\033[0m"
```

Usage:

```go
fmt.Println(Green + "Pet saved successfully!" + Reset)
```

`Reset` is important because it turns the terminal color back to normal.

---

## Troubleshooting

### `go: command not found`

Go is not installed or not added to your PATH.

Install Go from:

```bash
https://go.dev/dl/
```

Then restart your terminal and check:

```bash
go version
```

---

### `permission denied`

If you get permission issues while moving the executable:

```bash
mv gotermi /usr/local/bin/
```

Use:

```bash
sudo mv gotermi /usr/local/bin/
```

---

### `gotermi: command not found`

This means the executable is not installed globally or `/usr/local/bin` is not in your PATH.

Try:

```bash
echo $PATH
```

If `/usr/local/bin` is missing, add it to your shell config.

For zsh on Mac:

```bash
nano ~/.zshrc
```

Add this line:

```bash
export PATH="/usr/local/bin:$PATH"
```

Save and reload:

```bash
source ~/.zshrc
```

Then try:

```bash
gotermi
```

---

### Pet data is not resetting

Delete the saved JSON file:

```bash
rm ~/.gotermi/pet.json
```

Then run Gotermi again:

```bash
gotermi
```

---

### I changed the code but `gotermi` still runs the old version

Rebuild and reinstall:

```bash
go build -o gotermi
sudo mv gotermi /usr/local/bin/
```

Then run:

```bash
gotermi
```

---

## Tech Used

- Go
- JSON
- File handling
- Terminal ANSI colors
- CLI basics
- ASCII UI

---

## What This Project Demonstrates

This project shows:

- How to build a CLI app in Go
- How to use structs
- How to use pointers
- How to take terminal input
- How to create a menu loop
- How to save and load JSON data
- How to store app data in a hidden folder
- How to use terminal colors
- How to build a Go binary
- How to install a custom terminal command

---

## Final Result

After setup, users can run:

```bash
gotermi
```

from anywhere in their terminal and play with their own terminal pet.

Built with Go, JSON, terminal chaos, and emotional support from a tiny ASCII pet. 🐹🌸