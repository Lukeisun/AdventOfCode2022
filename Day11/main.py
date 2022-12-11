
import sys

class Monkey():
    starting_items = []
    operation = None
    modifier = None
    test = None
    tMonkey = None
    fMonkey = None
    def __init__(self, starting_items, operation, modifier, test, tMonkey, fMonkey ):
        self.starting_items = starting_items
        self.operation = operation
        self.modifier = modifier
        self.test = test
        self.tMonkey = tMonkey
        self.fMonkey = fMonkey
    def __str__(self):
        return "Monkey: " + str(self.starting_items) + " " + " " + str(self.modifier) + " " + str(self.test) + " " + str(self.tMonkey) + " " + str(self.fMonkey)
def add (a,b):
    return a+b
def multiply (a,b):
    return a * b
def isDivisible (a,b):
    return a % b == 0
if __name__ == "__main__":
    monkeys = []
    n = 0
    lcm = 1;
    for inp in sys.stdin:
        inp = inp.split(" ")
        if inp[0] == "Monkey":
            items = sys.stdin.readline().split(":")[1].split(",")
            items = [int(i.strip()) for i in items]
            op = sys.stdin.readline().split(":")[1].split(" ")
            try:
                modifier = int(op[5].strip())
            except:
                modifier = op[5].strip()
            if op[4] == "+":
                op = add
            elif op[4] == "*":
                op = multiply
            test = int(sys.stdin.readline().split(":")[1].split(" ")[3].strip())
            lcm *= test
            t = int(sys.stdin.readline().split(":")[1].split(" ")[4].strip())
            f = int(sys.stdin.readline().split(":")[1].split(" ")[4].strip())
            monkey = Monkey(items, op, modifier, test, t, f)
            monkeys.append(monkey)
    upLim = 20
    monkeyCount = [0] * len(monkeys)
    changedMonkeys = [i for i in range(0, len(monkeys))]
    for i in range(1, upLim+1):
        for j in range(0, len(monkeys)):
            for k in range(0, len(monkeys[j].starting_items)):
                monkeyCount[j] += 1
                currWorryLevel = monkeys[j].starting_items[k]
                modifier = monkeys[j].modifier
                if modifier == "old":
                    modifier = currWorryLevel
                currWorryLevel = monkeys[j].operation(currWorryLevel, modifier) # // 3 uncomment for part 1
                monkeys[j].starting_items[k] =  currWorryLevel  % lcm 
                flagForNextMonkey = False
                if isDivisible(currWorryLevel, monkeys[j].test):
                    flagForNextMonkey = True
                if flagForNextMonkey:
                    monkeys[monkeys[j].tMonkey].starting_items.append(monkeys[j].starting_items[k])
                else:
                    monkeys[monkeys[j].fMonkey].starting_items.append(monkeys[j].starting_items[k])
            monkeys[j].starting_items = []
        # print(f"ROUND {i}\n\n")
        # for j in range(0, len(monkeys)):
        #     print(monkeys[j], "\n")
    print(monkeyCount)
    firstMax = max(monkeyCount)
    monkeyCount.remove(firstMax)
    secondMax = max(monkeyCount)
    print(firstMax * secondMax)

