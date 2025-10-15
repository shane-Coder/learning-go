# Roadmap 2.0: The Path to Production-Grade Go

The first roadmap was about learning the fundamentals of the Go language. This second roadmap is focused on achieving mastery by learning the patterns, tools, and architectural principles required to build professional, production-ready applications.

---

## ✅ Phase 1: Deep Dive into Professional Testing (In Progress)

This phase is about building a safety net for our code. In the professional world, untested code is considered broken. Mastering testing enables confident refactoring and robust development.

- [x] **Table-Driven Tests**: The standard for writing clean, maintainable tests.
- [x] **Testing HTTP Handlers**: Learn to test API endpoints programmatically without a live server (`net/http/httptest`).
- [ ] **Mocking Dependencies**: Learn how to test business logic that depends on external systems (like databases) in isolation.
- [ ] **Benchmarking**: Learn to measure the speed of your code and compare the performance of different solutions.

---

## ⏳ Phase 2: Advanced Project Architecture (Up Next)

This phase focuses on structuring larger applications in a way that is scalable, maintainable, and easy to work on in a team.

- [ ] **Layered (Hexagonal) Architecture**: Learn to structure code into distinct layers (transport, service, repository) for a clean separation of concerns.
- [ ] **Dependency Injection**: Master the pattern of "injecting" dependencies (like your database) to make code more flexible and highly testable.
- [ ] **New Capstone Project**: Build a new, more advanced project (e.g., a "Multi-User API with Authentication") from the ground up using these professional patterns.

---

## 🚀 Phase 3: Performance and Optimization (Future Goal)

Once we can build and test robust applications, the final step is learning how to make them fast and efficient.

- [ ] **Profiling with `pprof`**: Learn to use Go's built-in profiler to find and fix CPU and memory bottlenecks in applications.
- [ ] **Advanced Concurrency Patterns**: Go beyond basic worker pools to learn more complex patterns for high-performance systems.