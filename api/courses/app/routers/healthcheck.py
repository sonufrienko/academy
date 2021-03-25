from fastapi import APIRouter
from datetime import datetime

server_start_time = datetime.today()
router = APIRouter()


@router.get("/healthcheck", status_code=200)
def read_healthcheck():
    up_time = datetime.today() - server_start_time
    return {
        "ok": True,
        "start_time": server_start_time.isoformat(),
        "up_time": round(up_time.total_seconds()),
    }