
import sys
def bfs(grid, endingPosition, m):
    queue = [endingPosition]
    dirArray = [(0,1), (0,-1), (1,0), (-1,0)]
    dirGrid = [['.' for i in range(len(grid[0]))] for j in range(len(grid))]
    while len(queue) > 0:
        curr = queue.pop(0)
        currRow = curr[0]
        currCol = curr[1]
        for dir in dirArray:
            nextRow = currRow + dir[0]
            nextCol = currCol + dir[1]
            if nextRow < 0 or nextCol < 0 or nextRow >= len(grid) or nextCol >= len(grid[0]):
                continue
            if grid[currRow][currCol] <= grid[nextRow][nextCol]+1 and (nextRow, nextCol) not in m:
                if dir == (0,1):
                    dirGrid[nextRow][nextCol] = '>'
                elif dir == (0,-1):
                    dirGrid[nextRow][nextCol] = '<'
                elif dir == (1,0):
                    dirGrid[nextRow][nextCol] = 'v'
                elif dir == (-1,0):
                    dirGrid[nextRow][nextCol] = '^'
                m[(nextRow, nextCol)] = curr[2] + 1
                queue.append((nextRow, nextCol, curr[2]+1)) 
    dirGrid[endingPosition[0]][endingPosition[1]] = 'E'
    dirGrid[20][0] = 'S'
    for row in range(len(dirGrid)):
        for col in range(len(dirGrid[row])):
            print(dirGrid[row][col], end='')
        print()
    return m
if __name__ == "__main__":
    grid =[]
    elevationsOfA = []
    for inp in sys.stdin:
        grid.append(list(inp.strip()))
    for row in range(len(grid)):
        for col in range(len(grid[row])):
            if grid[row][col] == "a":
                elevationsOfA.append((row, col))
            if grid[row][col] == "S":
                startingpos = (row, col, 1)
                grid[row][col] = 0
            elif grid[row][col] == "E":
                grid[row][col] = 25
                endingPosition = (row, col, 0)
            else:
                grid[row][col] = ord(grid[row][col]) - ord('a')
    m = dict(endingPosition=0)
    bfs(grid, endingPosition, m)
    print(m[(startingpos[0], startingpos[1])])
    min = float("inf")
    for a in elevationsOfA:
        try:
            if m[a] < min:
                min = m[a]
        except:
            continue
    print(min)
