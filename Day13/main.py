
import sys, json, copy

def buildlist(p):
    li = json.loads(p)
    return li

def inOrder(p1, p2):
    flag = True
    for i in range(len(p1)):
       # print(p1, p2)
        seq1 = p1.pop(0)
        try:
            seq2 = p2.pop(0)
        except:
            return False
        # print(seq1, seq2)
        isS1Int = isinstance(seq1, int)
        isS2Int = isinstance(seq2, int)
        isS1List = isinstance(seq1, list)
        isS2List = isinstance(seq2, list)
        if isS1Int and isS2Int:
            if seq1 == seq2: continue
            flag = seq1 < seq2
        if isS1Int and isS2List:
            seq1 = [seq1]
            isS1List = True
        if isS2Int and isS1List:
            seq2 = [seq2]
            isS2List = True
        if isS1List and isS2List:
            flag = inOrder(seq1, seq2)
        if flag is not None:
            return flag
    if len(p1) == 0 and len(p2) > 0:
        return True
if __name__ == "__main__":
    idx = 1
    count = 0
    sorted = []
    for p1 in sys.stdin:
        p2 = sys.stdin.readline()
        p1 = p1.strip()
        p2 = p2.strip()
        q1 = []
        q2 = []
        p1 = buildlist(p1)
        p2 = buildlist(p2)
        flag = inOrder(copy.deepcopy(p1), copy.deepcopy(p2))
        # print(flag)
        sorted.append(p1)
        sorted.append(p2)
        if flag:
            # print(idx)
            count += idx
        idx += 1
        sys.stdin.readline()
    sorted.append([[2]])
    sorted.append([[6]])
    for i in range(0, len(sorted)-1):
        for j in range(0, len(sorted) - i - 1):
            if not inOrder(copy.deepcopy(sorted[j]), copy.deepcopy(sorted[j+1])):
                temp1 = copy.deepcopy(sorted[j])
                sorted[j] = copy.deepcopy(sorted[j+1])
                sorted[j+1] = temp1
    #print(sorted)
    print((sorted.index([[2]]) + 1) * (sorted.index([[6]])+1))
    print(count)
    #print(inOrder([[1], 4], [1,[2,[3,[4,[5,6,0]]]],8,9]))