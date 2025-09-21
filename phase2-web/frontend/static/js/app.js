let currSimulate = null;
let currSteps = 0;
let simulateHistory = [];
let isRun = false;

let exampleSelect;
let tapeInput;
let maxStepsInput;
let runBtn;
let stepBtn;

let statusElement;
let stepsElement;
let tapeDisplay;
let historyList;

let jsonEditor;
let loadJsonBtn;

document.addEventListener('DOMContentLoaded', function() {
    initializeElements();
    setupEventListeners();
    loadDefaultExample();
});

function initializeElements() {
    exampleSelect = document.getElementById('example_select');
    tapeInput = document.getElementById('tape_input');
    maxStepsInput = document.getElementById('max_steps');
    runBtn = document.getElementById('run_btn');
    stepBtn = document.getElementById('step_btn');
    statusElement = document.getElementById('status');
    stepsElement = document.getElementById('steps');
    tapeDisplay = document.getElementById('tape_display');
    historyList = document.getElementById('history_list');
    jsonEditor = document.getElementById('json_editor');
    loadJsonBtn = document.getElementById('load_json_btn');
}

function setupEventListeners() {

    runBtn.addEventListener('click', runSimulation);
    stepBtn.addEventListener('click', stepThroughSimulation);
    loadJsonBtn.addEventListener('click', loadJsonConfiguration);
    exampleSelect.addEventListener('change', onExampleChange);
}

function loadDefaultExample() {

    tapeInput.value = "0011";
    maxStepsInput.value = "1000";
    const defaultExample = {
        "states": ["q0", "q1", "q_accept", "q_reject"],
        "input_alphabet": ["0", "1"],
        "tape_alphabet": ["0", "1", "X", "_"],
        "start_state": "q0",
        "accept_state": "q_accept",
        "reject_state": "q_reject",
        "transitions": {
            "q0,0": {"new_state": "q1", "write": "X", "direction": "R"},
            "q1,0": {"new_state": "q1", "write": "0", "direction": "R"},
            "q1,1": {"new_state": "q1", "write": "1", "direction": "R"},
            "q1,_": {"new_state": "q_accept", "write": "_", "direction": "N"}
        },
        "tape": "0011",
        "head_position": 0
    };
    
    jsonEditor.value = JSON.stringify(defaultExample, null, 2);
    updateStatus("Ready");
}


function updateStatus(message) {
    statusElement.textContent = message;
    statusElement.className = 'status';

    if (message === "Accepted") {
        statusElement.className = 'status accepted';
    } else if (message === "Rejected") {
        statusElement.className = 'status rejected';
    } else if (message === "Running...") {
        statusElement.className = 'status running';
    }
}

function onExampleChange() {
    const selectedExample = exampleSelect.value;
    
    if (selectedExample === "simple") {
        loadSimpleExample();
    } else if (selectedExample === "palindrome") {
        loadPalindromeExample();
    } else if (selectedExample === "binary") {
        loadBinaryExample();
    }
}

function loadSimpleExample() {
    const simpleExample = {
        "states": ["q0", "q1", "q_accept", "q_reject"],
        "input_alphabet": ["0", "1"],
        "tape_alphabet": ["0", "1", "X", "_"],
        "start_state": "q0",
        "accept_state": "q_accept",
        "reject_state": "q_reject",
        "transitions": {
            "q0,0": {"new_state": "q1", "write": "X", "direction": "R"},
            "q1,0": {"new_state": "q1", "write": "0", "direction": "R"},
            "q1,1": {"new_state": "q1", "write": "1", "direction": "R"},
            "q1,_": {"new_state": "q_accept", "write": "_", "direction": "N"}
        },
        "tape": "0011",
        "head_position": 0
    };
    
    tapeInput.value = "0011";
    jsonEditor.value = JSON.stringify(simpleExample, null, 2);
    updateStatus("Simple example loaded");
}

function loadPalindromeExample() {
    const palindromeExample = {
        "states": ["q0", "q1", "q2", "q3", "q4", "q_accept", "q_reject"],
        "input_alphabet": ["0", "1"],
        "tape_alphabet": ["0", "1", "X", "Y", "_"],
        "start_state": "q0",
        "accept_state": "q_accept",
        "reject_state": "q_reject",
        "transitions": {
            "q0,0": {"new_state": "q1", "write": "X", "direction": "R"},
            "q0,1": {"new_state": "q2", "write": "Y", "direction": "R"},
            "q0,_": {"new_state": "q_accept", "write": "_", "direction": "N"},
            "q1,0": {"new_state": "q1", "write": "0", "direction": "R"},
            "q1,1": {"new_state": "q1", "write": "1", "direction": "R"},
            "q1,_": {"new_state": "q3", "write": "_", "direction": "L"},
            "q2,0": {"new_state": "q2", "write": "0", "direction": "R"},
            "q2,1": {"new_state": "q2", "write": "1", "direction": "R"},
            "q2,_": {"new_state": "q4", "write": "_", "direction": "L"},
            "q3,X": {"new_state": "q_accept", "write": "X", "direction": "N"},
            "q3,Y": {"new_state": "q_reject", "write": "Y", "direction": "N"},
            "q4,X": {"new_state": "q_reject", "write": "X", "direction": "N"},
            "q4,Y": {"new_state": "q_accept", "write": "Y", "direction": "N"}
        },
        "tape": "1001",
        "head_position": 0
    };
    
    tapeInput.value = "1001";
    jsonEditor.value = JSON.stringify(palindromeExample, null, 2);
    updateStatus("Palindrome example loaded");
}

function loadBinaryExample() {
    const binaryExample = {
        "states": ["q0", "q1", "q2", "q3", "q_accept", "q_reject"],
        "input_alphabet": ["0", "1"],
        "tape_alphabet": ["0", "1", "_"],
        "start_state": "q0",
        "accept_state": "q_accept",
        "reject_state": "q_reject",
        "transitions": {
            "q0,0": {"new_state": "q0", "write": "0", "direction": "R"},
            "q0,1": {"new_state": "q0", "write": "1", "direction": "R"},
            "q0,_": {"new_state": "q1", "write": "_", "direction": "L"},
            "q1,0": {"new_state": "q2", "write": "1", "direction": "L"},
            "q1,1": {"new_state": "q1", "write": "0", "direction": "L"},
            "q2,0": {"new_state": "q2", "write": "0", "direction": "L"},
            "q2,1": {"new_state": "q2", "write": "1", "direction": "L"},
            "q2,_": {"new_state": "q_accept", "write": "_", "direction": "N"}
        },
        "tape": "101",
        "head_position": 0
    };
    
    tapeInput.value = "101";
    jsonEditor.value = JSON.stringify(binaryExample, null, 2);
    updateStatus("Binary counter example loaded");
}

function loadJsonConfiguration() {
    try {
        const jsonText = jsonEditor.value;
        const config = JSON.parse(jsonText);
        
        // Validate basic structure
        if (!config.states || !config.transitions || !config.tape) {
            throw new Error("Invalid JSON structure");
        }
        
        // Update tape input
        tapeInput.value = config.tape;
        
        updateStatus("JSON configuration loaded successfully");
    } catch (error) {
        updateStatus("Error loading JSON: " + error.message);
    }
}

function runSimulation() {
    try {
        // Get configuration from JSON editor
        const jsonText = jsonEditor.value;
        const config = JSON.parse(jsonText);
        
        // Update tape with user input
        config.tape = tapeInput.value;
        config.head_position = 0;
        
        // Get max steps
        const maxSteps = parseInt(maxStepsInput.value);
        
        // Prepare request
        const request = {
            definition: config,
            maxSteps: maxSteps
        };
        
        // Send to backend
        updateStatus("Running...");
        fetch('/api/simulate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request)
        })
        .then(response => response.json())
        .then(data => {
            currSimulate = data;
            currSteps = data.steps;
            simulateHistory = data.history || [];
            
            updateStatus(data.result);
            stepsElement.textContent = data.steps;
            displayTape(data.finalTape);
            displayHistory();
        })
        .catch(error => {
            updateStatus("Error: " + error.message);
        });
        
    } catch (error) {
        updateStatus("Error: " + error.message);
    }
}

function displayTape(tape) {
    if (!tape || !Array.isArray(tape)) {
        tapeDisplay.innerHTML = '<span class="error">Invalid tape data</span>';
        return;
    }
    
    let html = '<div class="tape-container">';
    tape.forEach((symbol, index) => {
        html += `<span class="tape-cell">${symbol}</span>`;
    });
    html += '</div>';
    
    tapeDisplay.innerHTML = html;
}

function displayHistory() {
    if (!simulateHistory || simulateHistory.length === 0) {
        historyList.innerHTML = '<div class="no-history">No execution history available</div>';
        return;
    }
    
    let html = '';
    simulateHistory.forEach((step, index) => {
        html += `<div class="history-step">
            <span class="step-number">Step ${step.step}</span>
            <span class="step-state">State: ${step.state}</span>
            <span class="step-head">Head: ${step.head}</span>
            <span class="step-action">${step.action}</span>
        </div>`;
    });
    
    historyList.innerHTML = html;
}

function stepThroughSimulation() {
    if (isRun) {
        updateStatus("Simulation already running");
        return;
    }
    
    try {
        // Get configuration from JSON editor
        const jsonText = jsonEditor.value;
        const config = JSON.parse(jsonText);
        
        // Update tape with user input
        config.tape = tapeInput.value;
        config.head_position = 0;
        
        // Get max steps
        const maxSteps = parseInt(maxStepsInput.value);
        
        // Prepare request
        const request = {
            definition: config,
            maxSteps: maxSteps
        };
        
        // Send to backend for step-by-step simulation
        updateStatus("Starting step-by-step simulation...");
        isRun = true;
        
        fetch('/api/simulate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request)
        })
        .then(response => response.json())
        .then(data => {
            currSimulate = data;
            currSteps = data.steps;
            simulateHistory = data.history || [];
            
            updateStatus("Step-by-step simulation completed");
            stepsElement.textContent = data.steps;
            displayTape(data.finalTape);
            displayHistory();
            isRun = false;
        })
        .catch(error => {
            updateStatus("Error: " + error.message);
            isRun = false;
        });
        
    } catch (error) {
        updateStatus("Error: " + error.message);
        isRun = false;
    }
}