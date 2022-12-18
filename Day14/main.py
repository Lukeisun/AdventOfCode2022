import sys, copy
def move(graph, sandX, sandY, moveDir, safePos):
    # print(sandX, sandY)
    if sandY +moveDir[0] >= len(graph) or sandX +moveDir[1] >= len(graph[0]):
        # graph[safePos[0]][safePos[1]] = "o"
        return False
    if graph[sandY + moveDir[0]][sandX + moveDir[1]] == ".":
        safePos = (sandY+moveDir[0], sandX+moveDir[1])
        return move(graph,sandX+moveDir[1], sandY+moveDir[0], [1,0], safePos)
    else:
        if moveDir == [1, 0]:
            moveDir = [1, -1]
        elif moveDir == [1, -1]:
            moveDir = [1, 1]
        else:
            # print(safePos)
            if graph[safePos[0]][safePos[1]] == "o":
                return False
            graph[safePos[0]][safePos[1]] = "o"
            return True
        return move(graph, sandX, sandY,moveDir, safePos)
    
if __name__ == "__main__":
    rocks = []
    minX = float("inf")
    minY = float("inf")
    maxX = float("-inf")
    maxY = float("-inf")
    for line in sys.stdin:
        pairs = line.split("->")
        rocks.append([])
        for pair in pairs:
            pair = pair.strip()
            pair = pair.split(",")
            p = (int(pair[0]), int(pair[1]))
            if p[0] > maxX:
                maxX = p[0]
            if p[0] < minX:
                minX = p[0]
            if p[1] > maxY:
                maxY = p[1]
            if p[1] < minY:
                minY = p[1]
            rocks[-1].append(p)
    # part 2
    graph = [["." for i in range(0, maxX-minX + 1000) ] for j in range(0, maxY+2+1)]
    for i in range(len(graph[-1])):
        graph[-1][i] = "#"
    # graph = [["." for i in range(0, maxX-minX + 1) ] for j in range(0, maxY+1)] # Uncomment for part 1
    minX = 0
    sandX = 500-minX
    graph[0][sandX] = "+"
    for rock in rocks:
        for idx, coords in enumerate(rock):
            x = coords[0]-minX 
            y = coords[1]
            graph[y][x] = "#"
            if idx + 1 < len(rock):
                x2 = rock[idx+1][0] - minX
                y2 = rock[idx+1][1] 
                smallerY = min(y, y2)
                smallerX = min(x, x2)
                if x == x2:
                    for i in range(abs(y-y2)):
                        graph[i+smallerY][x] ="#"
                if y == y2:
                    for i in range(abs(x-x2)):
                        graph[y][i+smallerX] ="#"
    i = 0
    while True:
        sandY = 0
        sandX = 500-minX
        moveDir = [1, 0]
        # prevGraph = copy.deepcopy(graph)
        flag = move(graph, sandX, sandY, moveDir, (sandY, sandX))
        if not flag:
            break
        # if (graph == prevGraph):
        #     break
        i +=1
    print(i)
    for i in range(0, len(graph)):
        for j in range(0, len(graph[i])):
            print(graph[i][j], end = "")
        print("")
