import uvicorn
from fastapi import FastAPI
from routers import courses, healthcheck

app = FastAPI()
app.include_router(courses.router, prefix="/v1")
app.include_router(healthcheck.router)


@app.on_event("startup")
def startup_event():
    # write logs
    pass


@app.on_event("shutdown")
def shutdown_event():
    # write logs
    pass


if __name__ == "__main__":
    uvicorn.run(app)
