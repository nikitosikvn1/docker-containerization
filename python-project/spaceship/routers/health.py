from enum import IntEnum

from fastapi import APIRouter
from pydantic import BaseModel

router = APIRouter()


class ServiceStatusCode(IntEnum):
    ok = 0
    bad = 1


class HealthcheckResult(BaseModel):
    service_status_code: ServiceStatusCode


@router.get('', response_model=HealthcheckResult)
def healthcheck() -> dict:
    return {'service_status_code': ServiceStatusCode.ok}
