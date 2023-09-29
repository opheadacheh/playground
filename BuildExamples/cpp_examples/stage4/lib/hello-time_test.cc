#include "gtest/gtest.h"
#include "lib/hello-time.h"

TEST(HelloTime, Add) {
  EXPECT_EQ(add(1, 2), 3);
}