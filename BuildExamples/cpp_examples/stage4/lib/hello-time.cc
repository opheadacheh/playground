#include "lib/hello-time.h"
#include <ctime>
#include <iostream>

void print_localtime() {
  std::time_t result = std::time(nullptr);
  std::cout << std::asctime(std::localtime(&result));
}

int add(int a, int b) {
  return a + b;
}