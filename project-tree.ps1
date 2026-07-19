# prerequisites, run this 2 line of code first below
# Set-ExecutionPolicy -Scope CurrentUser RemoteSigned
# Set-ExecutionPolicy -Scope Process Bypass

param(
    [int]$Depth = 999
)

if (-not (git rev-parse --is-inside-work-tree 2>$null)) {
    Write-Error "Not a git repository."
    exit 1
}

$files = git ls-files --cached --others --exclude-standard

if (-not $files) {
    Write-Host "."
    exit 0
}

$root = [ordered]@{}

foreach ($file in $files) {
    $parts = $file -split '[\\/]'
    $node = $root

    foreach ($part in $parts) {
        if (-not $node.Contains($part)) {
            $node[$part] = [ordered]@{}
        }

        $node = $node[$part]
    }
}

function Show-Tree {
    param(
        $Tree,
        [string]$Prefix = "",
        [int]$Level = 0
    )

    if ($Level -ge $Depth) {
        return
    }

    $keys = @($Tree.Keys) | Sort-Object

    for ($i = 0; $i -lt $keys.Count; $i++) {
        $key = $keys[$i]
        $isLast = ($i -eq $keys.Count - 1)

        if ($isLast) {
            Write-Host ($Prefix + "|--- " + $key)
            $next = $Prefix + "    "
        }
        else {
            Write-Host ($Prefix + "|--- " + $key)
            $next = $Prefix + "|   "
        }

        if ($Tree[$key].Count -gt 0) {
            Show-Tree -Tree $Tree[$key] -Prefix $next -Level ($Level + 1)
        }
    }
}

Write-Host "."
Show-Tree -Tree $root