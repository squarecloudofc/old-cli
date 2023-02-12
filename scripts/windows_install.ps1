#Requires -Version 5

$repository = "richaardev/squarecloud-cli" # just for maybe future changes
$osArchitecture = if ([Environment]::Is64BitProcess) { "amd64" } else { "386" }
$downloadUri = (
  Invoke-WebRequest -Uri "https://api.github.com/repos/$repository/releases/latest" -UseBasicParsing | 
  select-string -Pattern "https://github.com/$repository/releases/download/v\d+\.\d+\.\d+/squarecloud_\d+\.\d+\.\d+_windows_$osArchitecture.zip"
  ).Matches.Value
$tempZip = New-TemporaryFile | Rename-Item -NewName { $_ -replace 'tmp', 'zip' } -PassThru
$cliPath = Join-Path "$env:LOCALAPPDATA" "Square Cloud CLI"
$cliBinPath = Join-Path "$cliPath" "bin"

Write-Host ""
Write-Host "Installing Square Cloud CLI...`n"

Invoke-WebRequest -Uri $downloadUri -UseBasicParsing -OutFile $tempZip | Out-Null
Expand-Archive -Path $tempZip -DestinationPath $cliBinPath -Force | Out-Null

Remove-Item $tempZip -Force

# add cli to Path
if([Environment]::GetEnvironmentVariable("PATH", "User") -split ";" -cnotcontains $cliBinPath) {
  $env:Path += $cliBinPath
  [Environment]::SetEnvironmentVariable(
    "Path", 
    [Environment]::GetEnvironmentVariable("PATH", "User") + [IO.Path]::PathSeparator + $cliBinPath,
    "User"
  )
} 


Write-Host "[$([char]8730)] Installation Successfuly!" -ForegroundColor Green
Write-Host ""
Write-Host "To run Square Cloud CLI, execute:" -ForegroundColor Cyan
Write-Host "  squarecloud --help"
Write-Host ""
Write-Host "(May you need to restart your terminal)" -ForegroundColor DarkGray
Write-Host ""