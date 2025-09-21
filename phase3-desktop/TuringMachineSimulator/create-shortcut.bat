@echo off 
echo Creating desktop shortcut... 
powershell "$WshShell = New-Object -comObject WScript.Shell; $Shortcut = $WshShell.CreateShortcut('%USERPROFILE%\Desktop\Turing Machine Simulator.lnk'); $Shortcut.TargetPath = '%~dp0TuringMachineSimulator.exe'; $Shortcut.Save()" 
echo Shortcut created on desktop! 
pause 
