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