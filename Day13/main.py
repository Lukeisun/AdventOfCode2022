import sys, re
def buildlist(p):
    li = list()
    idx = 0
    for i in range(1, len(p)):
        print(li, idx)
        if p[i] == "[":
            idx += 1
            li.append(list())
        elif p[i] == "]":
            idx -= 1
        elif p1[i] == ",":
            continue
        else:
            if isinstance(li[idx], list):
                li[idx].append(p[i])
            else:
                li.append(p[i])
    return li 
if __name__ == "__main__":
    for p1 in sys.stdin:
        p2 = sys.stdin.readline()
        p1 = p1.strip()
        p2 = p2.strip()
        q1 = []
        q2 = []
        print(buildlist(p1))
        # print(p2)
        # print(buildlist(p1))
        #print(p1, "\n", p2)
        sys.stdin.readline()