#include <iostream>
using namespace std;
#include <iostream>
#include <fstream>
#include <string>
#include <queue>

int main() {
  std::ifstream file ("input.in");
  string line;
  if (file.is_open()) {
    int currCount = 0;
    priority_queue<int> top;
    while (std::getline(file, line)) {
        if (line.empty()) {
          top.push(currCount);
          currCount = 0;
          continue;
        }
        currCount += std::stoi(line);
      }
      int sum = 0;
      for (int i = 0; i < 3; i++) {
        std::cout << top.top() << std::endl;
        sum += top.top();
        top.pop();
      }
      std::cout <<  sum << '\n';
    }
    return 0;
}