import sys

if __name__ == "__main__":
    sensors = []
    beacons = []
    for line in sys.stdin:
        halves = line.split(": ")
        lhs = halves[0].split(" ")
        rhs = halves[1].split(" ")
        s_x = int((lhs[2].split("="))[1].replace(",", ""))
        s_y = int(lhs[3].split("=")[1])
        b_x = int(rhs[4].split("=")[1].replace(",", ""))
        b_y = int(rhs[5].split("=")[1])
        print(s_x, s_y, b_x, b_y)