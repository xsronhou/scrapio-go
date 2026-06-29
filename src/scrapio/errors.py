from __future__ import annotations
from typing import Any


class ApiError(Exception):
    status_code: int
    request_id: str
    code: str

    def __init__(self, status_code: int, body: dict[str, Any]) -> None:
        error = body.get("error", {})
        super().__init__(error.get("message", "Unknown error"))
        self.status_code = status_code
        self.request_id = body.get("request_id", "")
        self.code = error.get("code", "unknown")


class AuthError(ApiError):
    def __init__(self, body: dict[str, Any]) -> None:
        super().__init__(401, body)


class RateLimitError(ApiError):
    def __init__(self, body: dict[str, Any]) -> None:
        super().__init__(429, body)


class CreditsExhaustedError(ApiError):
    def __init__(self, body: dict[str, Any]) -> None:
        super().__init__(402, body)
