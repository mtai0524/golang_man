# Svelte Desktop App powered by Wails 🚀

A modern, high-performance desktop application built with the perfect blend of **Go**'s backend power and **Svelte**'s seamless frontend experience.

## 🛠 Technology Stack

This project utilizes a modern layered architecture to optimize both developer productivity and application performance:

- **[Wails v2](https://wails.io/):** A framework for building desktop applications using Go and web technologies. Unlike bulky Electron apps, Wails uses the OS's native rendering engine (WebView2 on Windows, WebKit on macOS/Linux), resulting in an extremely lightweight application with minimal memory footprint.
- **[Go (Golang)](https://go.dev/):** Serves as the core logic handler, interacting with the operating system (file system, local database, network, etc.) and executing high-performance tasks on the backend.
- **[Svelte](https://svelte.dev/):** A frontend framework that provides an outstanding interactive UI experience. It compiles code directly into highly optimized vanilla JavaScript without using a Virtual DOM, leading to superior rendering performance.
- **[Vite](https://vitejs.dev/):** A next-generation frontend bundler that delivers lightning-fast startup times and smooth Hot Module Replacement (HMR).

## 📦 System Requirements

Before you begin, ensure your system has the following environment set up:
- [Go](https://go.dev/dl/) (version 1.18 or higher)
- [Node.js](https://nodejs.org/) (version 14+)
- [npm](https://www.npmjs.com/) (usually bundled with Node.js)

Install the Wails CLI on your system:
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## 🚀 Project Initialization

To initialize a brand new project using the Svelte template, run the following command:
```bash
wails init -n your-project-name -t svelte
```

## 💻 Live Development

To develop the application with Hot Reload for the Frontend (Vite) and auto-recompilation for the Backend (Go) when code changes occur:

```bash
# Navigate to the project directory
cd svelte-desktop

# Start development mode
wails dev
```

**💡 Pro Tip: Debugging with a Web Browser**

By default, `wails dev` opens a native desktop window. However, debugging the UI and inspecting elements on a native window can sometimes be inconvenient. 
To use the powerful DevTools of Chrome/Edge while still being able to call native Go functions:
- Wails provides a linked Dev Server (IPC Proxy) running at: `http://localhost:34115`
- Open this URL in your favorite web browser. Buttons and Go function calls will work exactly as they do in the Desktop App!

## 🏗 Building for Production

When your application is complete and ready to be distributed as a standalone executable file (e.g., `.exe` on Windows):

```bash
wails build
```

- Behind the scenes, Wails will automatically compile the Svelte code, embed it locally into the Go source code, and build it into a single file.
- The output executable file (`svelte-desktop.exe`) will be ready in the `build/bin/` directory.

---
*Built with ❤️ using Wails and Svelte.*
