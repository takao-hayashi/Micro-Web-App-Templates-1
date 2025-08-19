# Micro Web Apps – 8 Tiny Apps in JS / Python / Go / Rust

> 🧩 A collection of 8 ultra-simple web apps & APIs.  
> Clone → run → screenshot-ready(well no but...)
>  
> Author: **Takao Hayashi**  


## 📦 Apps

### JavaScript (Frontend only)
- **QR Generator** – Generate QR codes instantly in browser
- **Markdown Preview** – Live preview for Markdown

### Python (FastAPI)
- **Slugify API** – Convert text → clean URL slugs
- **Tiny TODO API** – Minimal in-memory task manager

### Go (Standard Library)
- **URL Shortener** – In-memory short links
- **JSON Pretty** – Beautify JSON via POST

### Rust (Warp)
- **UUID API** – Generate random UUID v4
- **Echo API** – POST JSON, get it echoed back


## 🚀 Quickstart

Clone the repo:

```bash
git clone https://github.com/<your-username>/micro-web-apps.git
cd micro-web-apps


⸻

🟨 JavaScript

No build needed. Just open in browser:

open js/qr-gen/index.html
open js/md-preview/index.html


⸻

🐍 Python

Requires Python 3.10+

cd python/fastapi-slug
pip install -r requirements.txt
python main.py
# -> http://localhost:8000/slug

cd ../fastapi-todo
pip install -r requirements.txt
python main.py
# -> http://localhost:8001/items


⸻

🟦 Go

Requires Go 1.22+

cd go/url-shortener
go run .
# -> POST http://localhost:8080/create {"url":"https://example.com"}

cd ../json-pretty
go run .
# -> POST http://localhost:8081/pretty {"a":1}


⸻

🦀 Rust (Warp)

Requires Rust 1.80+

cd rust
cargo run -p warp-uuid
# -> GET http://localhost:3000/uuid

cargo run -p warp-echo
# -> POST http://localhost:3001/echo {"msg":"hello"}


⸻

📸 Screenshots

Sorry that I did not take any Gifs. But I think all the codes are able to work...

⸻

## 📅 Weekly Template Release
This repository will grow with new template every week!
Stay tuned for fresh ideas in multiple languages!

### ✅ Completed
- JavaScript → QR Generator / Markdown Preview
- Python → Slugify API / Tiny TODO API
- Go → URL Shortener / JSON Pretty
- Rust → UUID API / Echo API

### 🔮 Next Week’s Preview
(Upcoming): 
TypeScript (Node.js + Express) templates
  - JWT Auth API
  - File Upload API
Ruby (Sinatra) mini apps  
PHP (Laravel/Lumen) tiny APIs  
C# (ASP.NET Core Minimal APIs) 
Deno/Bun fun web tools  

> 🚀 If you want to suggest the next template, open an Issue!
⸻

## 🌐 Live Demo (Example Sites)
_These are not this repo’s deployments, but example sites showing similar functionality.

- QR Generator → [qr-code-generator.com](https://www.qr-code-generator.com/)  
- Markdown Preview → [dillinger.io](https://dillinger.io/)  
- URL Shortener → [bitly.com](https://bitly.com/)  
- JSON Pretty → [jsonformatter.org](https://jsonformatter.org/json-pretty-print)  
- UUID API → [uuidtools.com/api](https://www.uuidtools.com/api/generate)  
- Echo API → [postman-echo.com](https://postman-echo.com/post)  

⸻
📜 License

MIT License © 2025 Takao Hayashi