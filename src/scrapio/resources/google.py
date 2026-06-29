from __future__ import annotations
from .._http import SyncHttpClient, AsyncHttpClient
from ..types import GoogleSearchParams, GoogleSearchResponse


class GoogleResource:
    def __init__(self, http: SyncHttpClient) -> None:
        self._http = http

    def search(self, params: GoogleSearchParams) -> GoogleSearchResponse:
        return self._http.request(
            "GET", "/v1/google/search",
            params=params.model_dump(exclude_none=True),
            response_model=GoogleSearchResponse,
        )


class AsyncGoogleResource:
    def __init__(self, http: AsyncHttpClient) -> None:
        self._http = http

    async def search(self, params: GoogleSearchParams) -> GoogleSearchResponse:
        return await self._http.request(
            "GET", "/v1/google/search",
            params=params.model_dump(exclude_none=True),
            response_model=GoogleSearchResponse,
        )
