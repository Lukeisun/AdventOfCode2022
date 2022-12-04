
import numpy as np
if __name__ == "__main__":
    print('hello')
    inp = input()
    count = 0
    count_2 = 0
    while True:
        print(inp)
        set1 = set({})
        set2 = set({})
        arr = inp.split(',');
        x1, y1 = arr[0].split('-')
        x2, y2 = arr[1].split('-')
        for i in range(int(x1), int(y1)+1):
            set1.add(i)
        for i in range(int(x2), int(y2)+1):
            set2.add(i)
        if(len(set1) != len(set2)):
            largerSet = set1 if len(set1) > len(set2) else set2
            smallerSet = set1 if len(set1) < len(set2) else set2
        else:
            largerSet = set1
            smallerSet = set2
        flag = True
        for i in smallerSet:
            if i not in largerSet:
                flag = False
                break
        # part 2
        for i in smallerSet:
            if i in largerSet:
                count_2 += 1
                break
        if flag:
            count += 1
        try:
            inp = input()
        except Exception as e:
            break
    print(count)
    print(count_2)