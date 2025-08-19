from fastapi import FastAPI
from pydantic import BaseModel
import re, uvicorn

app = FastAPI(title="Slugify API")

class In(BaseModel):
    text: str

def slugify(s: str) -> str:
    s = s.lower()
    s = re.sub(r'[^a-z0-9\s-]', '', s)
    s = re.sub(r'\s+', '-', s).strip('-')
    s = re.sub(r'-+', '-', s)
    return s

@app.post("/slug")
def make_slug(body: In): return {"slug": slugify(body.text)}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)