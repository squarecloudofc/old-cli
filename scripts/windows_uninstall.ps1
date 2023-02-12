#Requires -Version 5


$cliPath = Join-Path "$env:LOCALAPPDATA" "Square Cloud CLI"
$cliBinPath = Join-Path "$cliPath" "bin"

Write-Host ""
Write-Host "Uninstalling Square Cloud CLI...`n"

Remove-Item -LiteralPath $cliPath -Force -Recurse | Out-Null

# remove cli from Path
if ([Environment]::GetEnvironmentVariable("PATH", "User") -split ";" -ccontains $cliBinPath) {
  [Environment]::SetEnvironmentVariable(
    'PATH',
    ([Environment]::GetEnvironmentVariable(
      'PATH',
      'User'
    ).Split(';') | Where-Object { $_ -ne $cliBinPath }) -join ';',
    'User'
  )
} 

Write-Host ""
Write-Host "[$([char]8730)] Uninstall completed successfully!`n" -ForegroundColor Green
