from fastapi import FastAPI
from pydantic import BaseModel
from typing import List
import uvicorn, uuid

app = FastAPI(title="Tiny TODO")
DB = {}  # {id: {"id":..., "text":..., "done":bool}}

class ItemIn(BaseModel): text: str
class ItemOut(ItemIn): id: str; done: bool = False

@app.get("/items", response_model=List[ItemOut])
def list_items(): return list(DB.values())

@app.post("/items", response_model=ItemOut)
def add_item(it: ItemIn):
    i = {"id": uuid.uuid4().hex[:8], "text": it.text, "done": False}
    DB[i["id"]] = i; return i

@app.post("/items/{iid}/toggle", response_model=ItemOut)
def toggle(iid: str):
    DB[iid]["done"] = not DB[iid]["done"]; return DB[iid]

if __name__ == "__main__": uvicorn.run(app, host="0.0.0.0", port=8001)