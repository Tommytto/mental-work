cmake_minimum_required(VERSION 3.20)
project(cpp_practice LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# ---------------------------
# 1) Your normal executable(s)
# ---------------------------
add_executable(binsearch
  binsearch/main.cpp
)

# If later you want include dirs:
# target_include_directories(binsearch PRIVATE some_folder)

# ---------------------------
# 2) GoogleTest + tests
# ---------------------------
include(FetchContent)

FetchContent_Declare(
  googletest
  GIT_REPOSITORY https://github.com/google/googletest.git
  GIT_TAG v1.17.0
)

# For macOS / clang this avoids issues with gtest and shared CRT settings (mostly Windows-related,
# but leaving it is common and harmless)
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)

FetchContent_MakeAvailable(googletest)

enable_testing()

add_executable(binsearch_tests
  binsearch/tests.cpp
)

target_link_libraries(binsearch_tests
  PRIVATE
  GTest::gtest_main
)

include(GoogleTest)
gtest_discover_tests(binsearch_tests)