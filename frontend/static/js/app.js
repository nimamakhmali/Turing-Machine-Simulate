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

document.Add_Event_Listener('DOMContentLoaded', function() {
    initializeElements();
    setupEventListeners();
    loadDefaultExample();
});

function initial_Elements() {
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