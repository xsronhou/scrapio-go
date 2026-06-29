from __future__ import annotations

import asyncio
import time
from typing import Any, Optional, Type, TypeVar

import httpx
from pydantic import BaseModel

from .errors import ApiError, AuthError, CreditsExhaustedError, RateLimitError

T = TypeVar("T", bound=BaseModel)

RETRYABLE_STATUS = {429, 503}
DEFAULT_TIMEOUT = 30.0
DEFAULT_MAX_RETRIES = 3


def _raise_for_status(status_code: int, body: dict[str, Any]) -> None:
    code = body.get("error", {}).get("code", "")
    if status_code == 401:
        raise AuthError(body)
    if status_code == 429:
        raise RateLimitError(body)
    if status_code == 402 or code == "credits_exhausted":
        raise CreditsExhaustedError(body)
    raise ApiError(status_code, body)


class SyncHttpClient:
    def __init__(
        self,
        base_url: str,
        api_key: str,
        timeout: float = DEFAULT_TIMEOUT,
        max_retries: int = DEFAULT_MAX_RETRIES,
    ) -> None:
        self._base_url = base_url.rstrip("/")
        self._headers = {"Authorization": f"Bearer {api_key}"}
        self._timeout = timeout
        self._max_retries = max_retries
        self._client = httpx.Client(
            base_url=self._base_url,
            headers=self._headers,
            timeout=self._timeout,
        )

    def close(self) -> None:
        self._client.close()

    def __enter__(self) -> "SyncHttpClient":
        return self

    def __exit__(self, *args: Any) -> None:
        self.close()

    def request(
        self,
        method: str,
        path: str,
        *,
        params: Optional[dict[str, Any]] = None,
        json: Optional[Any] = None,
        response_model: Type[T],
        timeout: Optional[float] = None,
    ) -> T:
        clean_params = {k: v for k, v in (params or {}).items() if v is not None}

        for attempt in range(self._max_retries + 1):
            res = self._client.request(
                method,
                path,
                params=clean_params or None,
                json=json,
                timeout=timeout or self._timeout,
            )
            if res.is_success:
                return response_model.model_validate(res.json())

            body: dict[str, Any] = {}
            try:
                body = res.json()
            except Exception:
                body = {"request_id": "", "error": {"code": "unknown", "message": res.text}}

            if res.status_code in RETRYABLE_STATUS and attempt < self._max_retries:
                backoff = min(1.0 * (2**attempt), 8.0)
                time.sleep(backoff)
                continue

            _raise_for_status(res.status_code, body)

        raise RuntimeError("Unexpected end of retry loop")  # unreachable


class AsyncHttpClient:
    def __init__(
        self,
        base_url: str,
        api_key: str,
        timeout: float = DEFAULT_TIMEOUT,
        max_retries: int = DEFAULT_MAX_RETRIES,
    ) -> None:
        self._base_url = base_url.rstrip("/")
        self._headers = {"Authorization": f"Bearer {api_key}"}
        self._timeout = timeout
        self._max_retries = max_retries
        self._client = httpx.AsyncClient(
            base_url=self._base_url,
            headers=self._headers,
            timeout=self._timeout,
        )

    async def aclose(self) -> None:
        await self._client.aclose()

    async def __aenter__(self) -> "AsyncHttpClient":
        return self

    async def __aexit__(self, *args: Any) -> None:
        await self.aclose()

    async def request(
        self,
        method: str,
        path: str,
        *,
        params: Optional[dict[str, Any]] = None,
        json: Optional[Any] = None,
        response_model: Type[T],
        timeout: Optional[float] = None,
    ) -> T:
        clean_params = {k: v for k, v in (params or {}).items() if v is not None}

        for attempt in range(self._max_retries + 1):
            res = await self._client.request(
                method,
                path,
                params=clean_params or None,
                json=json,
                timeout=timeout or self._timeout,
            )
            if res.is_success:
                return response_model.model_validate(res.json())

            body: dict[str, Any] = {}
            try:
                body = res.json()
            except Exception:
                body = {"request_id": "", "error": {"code": "unknown", "message": res.text}}

            if res.status_code in RETRYABLE_STATUS and attempt < self._max_retries:
                backoff = min(1.0 * (2**attempt), 8.0)
                await asyncio.sleep(backoff)
                continue

            _raise_for_status(res.status_code, body)

        raise RuntimeError("Unexpected end of retry loop")  # unreachable
