from __future__ import annotations
from .._http import SyncHttpClient, AsyncHttpClient
from ..types import FetchRequest, FetchResponse


class FetchResource:
    def __init__(self, http: SyncHttpClient) -> None:
        self._http = http

    def fetch(self, request: FetchRequest, *, timeout: float | None = None) -> FetchResponse:
        return self._http.request(
            "POST", "/v1/fetch",
            json=request.model_dump(exclude_none=True),
            response_model=FetchResponse,
            timeout=timeout,
        )


class AsyncFetchResource:
    def __init__(self, http: AsyncHttpClient) -> None:
        self._http = http

    async def fetch(self, request: FetchRequest, *, timeout: float | None = None) -> FetchResponse:
        return await self._http.request(
            "POST", "/v1/fetch",
            json=request.model_dump(exclude_none=True),
            response_model=FetchResponse,
            timeout=timeout,
        )
