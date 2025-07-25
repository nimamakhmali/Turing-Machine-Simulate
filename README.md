# Turing-Machine-Go

# 🧠 Turing Machine Simulator in Golang

A formal, academic-grade simulator for deterministic Turing machines, written in Go.

This project implements a general-purpose **Turing Machine (TM) simulator** capable of parsing machine definitions and simulating their step-by-step execution. It is built with educational clarity and structural rigor in mind, aimed at illustrating the principles of computability theory and formal models of computation.

---

## 📘 Project Description

This simulator accepts a Turing Machine definition—comprising states, symbols, transitions, and an initial tape configuration—and runs the computation based on deterministic transition rules.

The goal is to bridge theoretical foundations of computation with real-world programming by:

- Translating abstract machine definitions into concrete simulation
- Reinforcing the Church–Turing thesis via a programmable model
- Demonstrating how high-level languages like Go can simulate universal computation

> 🧩 **Applications**: This simulator can be used for educational demonstrations, testing custom TMs, or validating formal properties of languages defined over TMs.

---

## 💡 Features

- Accepts TM definitions in clean JSON format
- Supports arbitrary alphabets and tape sizes
- Handles right (`R`), left (`L`), and neutral (`N`) head movements
- Detects acceptance, rejection, and halting conditions
- Customizable initial tape and head position
- Step-limited simulation to avoid infinite loops

---

## 📂 Project Structure



---

## 🔧 How to Run

1. Clone the repository:
```bash
git clone https://github.com/nimamakhmali/turing-machine-go.git
cd turing-machine-go
```
📄 Sample Input Format

```bash
{
  "states": ["q0", "q1", "q_accept", "q_reject"],
  "input_alphabet": ["0", "1"],
  "tape_alphabet": ["0", "1", "X", "_"],
  "start_state": "q0",
  "accept_state": "q_accept",
  "reject_state": "q_reject",
  "transitions": {
    "q0,0": { "new_state": "q1", "write": "X", "direction": "R" },
    "q1,1": { "new_state": "q1", "write": "1", "direction": "R" },
    "q1,_": { "new_state": "q_accept", "write": "_", "direction": "N" }
  },
  "tape": "0011",
  "head_position": 0
}

```

📚 Theoretical Context

This project directly implements concepts from the Theory of Computation, particularly:

    Deterministic Turing Machines (DTM)

    Formal languages and automata

    Acceptance/rejection conditions

    Church–Turing thesis (simulation of computation)

Such simulations illustrate that any computable function can be modeled algorithmically and verified symbolically—core ideas in theoretical computer science
