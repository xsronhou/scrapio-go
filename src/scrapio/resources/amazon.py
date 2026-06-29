from __future__ import annotations
from typing import Optional
from .._http import SyncHttpClient, AsyncHttpClient
from ..types import AmazonProductResponse, AmazonSearchResponse


class AmazonResource:
    def __init__(self, http: SyncHttpClient) -> None:
        self._http = http

    def get_product(self, asin: str, *, country: Optional[str] = None) -> AmazonProductResponse:
        return self._http.request(
            "GET", "/v1/amazon/product",
            params={"asin": asin, "country": country},
            response_model=AmazonProductResponse,
        )

    def search(self, query: str, *, country: Optional[str] = None, page: Optional[int] = None) -> AmazonSearchResponse:
        return self._http.request(
            "GET", "/v1/amazon/search",
            params={"query": query, "country": country, "page": page},
            response_model=AmazonSearchResponse,
        )

    def queue_search_crawl(self, query: str, *, country: Optional[str] = None) -> dict:
        from pydantic import RootModel
        from typing import Any
        class _R(RootModel[dict[str, Any]]): pass
        result = self._http.request(
            "GET", "/v1/amazon/search/crawl",
            params={"query": query, "country": country},
            response_model=_R,
        )
        return result.root


class AsyncAmazonResource:
    def __init__(self, http: AsyncHttpClient) -> None:
        self._http = http

    async def get_product(self, asin: str, *, country: Optional[str] = None) -> AmazonProductResponse:
        return await self._http.request(
            "GET", "/v1/amazon/product",
            params={"asin": asin, "country": country},
            response_model=AmazonProductResponse,
        )

    async def search(self, query: str, *, country: Optional[str] = None, page: Optional[int] = None) -> AmazonSearchResponse:
        return await self._http.request(
            "GET", "/v1/amazon/search",
            params={"query": query, "country": country, "page": page},
            response_model=AmazonSearchResponse,
        )
