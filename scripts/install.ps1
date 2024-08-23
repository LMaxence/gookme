# Define variables
$repoOwner = "LMaxence"
$repoName = "gookme"
$binaryName = "gookme"
$version = "latest"

# Determine the architecture
$arch = if ([System.Environment]::Is64BitOperatingSystem) { "amd64" } else { "arm64" }

# Construct the download URL
$url = "https://github.com/$repoOwner/$repoName/releases/$version/download/$binaryName-windows-$arch.exe"

# Define the destination path
$destinationPath = "$env:USERPROFILE\Downloads\$binaryName.exe"

# Download the binary
Write-Host "Downloading $binaryName from $url..."
Invoke-WebRequest -Uri $url -OutFile $destinationPath

# Make the binary executable (not necessary on Windows, but we'll move it to a PATH directory)
$installPath = "$env:ProgramFiles\$binaryName"
if (-Not (Test-Path -Path $installPath)) {
    New-Item -ItemType Directory -Path $installPath
}
Move-Item -Path $destinationPath -Destination "$installPath\$binaryName.exe"

# Add the install path to the system PATH if not already present
$path = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
if ($path -notlike "*$installPath*") {
    [System.Environment]::SetEnvironmentVariable("Path", "$path;$installPath", [System.EnvironmentVariableTarget]::Machine)
    Write-Host "Added $installPath to the system PATH."
}

Write-Host "$binaryName installed successfully."