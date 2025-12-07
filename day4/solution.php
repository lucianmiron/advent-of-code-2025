<?php

$grid = file('input.txt', FILE_IGNORE_NEW_LINES);

$rows = count($grid);
$cols = $rows > 0 ? strlen($grid[0]) : 0;

$directions = [
    [-1, -1], [-1, 0], [-1, 1],
    [0, -1],          [0, 1],
    [1, -1],  [1, 0], [1, 1],
];

$accessibleCount = 0;

for ($r = 0; $r < $rows; $r++) {
    for ($c = 0; $c < $cols; $c++) {
        if ($grid[$r][$c] !== '@') {
            continue;
        }

        $adjacentRolls = 0;
        foreach ($directions as $dir) {
            $nr = $r + $dir[0];
            $nc = $c + $dir[1];
            if ($nr >= 0 && $nr < $rows && $nc >= 0 && $nc < $cols && $grid[$nr][$nc] === '@') {
                $adjacentRolls++;
            }
        }

        if ($adjacentRolls < 4) {
            $accessibleCount++;
        }
    }
}

echo "Accessible rolls: $accessibleCount\n";
