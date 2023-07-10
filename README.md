# README

Aimai-moko (fuzzy) finder for personal use in pwsh.

On Japanese PCs, the default of `[System.Console]::OutputEncoding` is not `UTF-8`, so it cannot properly receive multibyte characters in standard output.

Therefore, when dealing with multibyte characters, specify the `--bytearr` option to convert the result into a byte array for output, and decode it as UTF-8 with the following PowerShell function.

```powershell
function mokof {
    param(
        [switch]$ascii
    )
    $exePath = "PATH\TO\mokof.exe"
    if ($ascii) {
        $input | & $exePath | Write-Output
    }
    else {
        $byte = $input | & $exePath "--bytearr"
        if ($LASTEXITCODE -eq 0) {
            [System.Text.Encoding]::UTF8.GetString($byte) | Write-Output
        }
    }
}
```
