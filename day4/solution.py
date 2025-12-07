def solve(input_file):
    with open(input_file) as f:
        grid = [line.rstrip('\n') for line in f]

    rows = len(grid)
    cols = len(grid[0]) if rows > 0 else 0

    directions = [
        (-1, -1), (-1, 0), (-1, 1),
        (0, -1),          (0, 1),
        (1, -1),  (1, 0), (1, 1)
    ]

    accessible_count = 0

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] != '@':
                continue

            adjacent_rolls = 0
            for dr, dc in directions:
                nr, nc = r + dr, c + dc
                if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == '@':
                    adjacent_rolls += 1

            if adjacent_rolls < 4:
                accessible_count += 1

    return accessible_count


if __name__ == "__main__":
    result = solve("input.txt")
    print(f"Accessible rolls: {result}")
