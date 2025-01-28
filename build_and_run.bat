@echo off

:: Build the Go project
echo Building the project...
go build -o web-page-analyzer.exe

:: Check if the build was successful
if %ERRORLEVEL% == 0 (
    echo Build successful!
    :: Run the project
    echo Running the project...
    web-page-analyzer.exe
) else (
    echo Build failed!
)
