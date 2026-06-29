from __future__ import annotations
from typing import Optional
from .._http import SyncHttpClient, AsyncHttpClient
from ..types import WalmartProductResponse, WalmartSearchResponse


class WalmartResource:
    def __init__(self, http: SyncHttpClient) -> None:
        self._http = http

    def get_product(self, product_id: str, *, country: Optional[str] = None) -> WalmartProductResponse:
        return self._http.request(
            "GET", "/v1/walmart/product",
            params={"product_id": product_id, "country": country},
            response_model=WalmartProductResponse,
        )

    def search(self, query: str, *, country: Optional[str] = None, page: Optional[int] = None) -> WalmartSearchResponse:
        return self._http.request(
            "GET", "/v1/walmart/search",
            params={"query": query, "country": country, "page": page},
            response_model=WalmartSearchResponse,
        )

    def queue_search_crawl(self, query: str, *, country: Optional[str] = None) -> dict:
        from pydantic import RootModel
        from typing import Any
        class _R(RootModel[dict[str, Any]]): pass
        result = self._http.request(
            "GET", "/v1/walmart/search/crawl",
            params={"query": query, "country": country},
            response_model=_R,
        )
        return result.root


class AsyncWalmartResource:
    def __init__(self, http: AsyncHttpClient) -> None:
        self._http = http

    async def get_product(self, product_id: str, *, country: Optional[str] = None) -> WalmartProductResponse:
        return await self._http.request(
            "GET", "/v1/walmart/product",
            params={"product_id": product_id, "country": country},
            response_model=WalmartProductResponse,
        )

    async def search(self, query: str, *, country: Optional[str] = None, page: Optional[int] = None) -> WalmartSearchResponse:
        return await self._http.request(
            "GET", "/v1/walmart/search",
            params={"query": query, "country": country, "page": page},
            response_model=WalmartSearchResponse,
        )
