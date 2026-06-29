from __future__ import annotations
from typing import Optional
from ._http import SyncHttpClient, AsyncHttpClient, DEFAULT_TIMEOUT, DEFAULT_MAX_RETRIES
from .resources.fetch import FetchResource, AsyncFetchResource
from .resources.jobs import JobsResource, AsyncJobsResource
from .resources.google import GoogleResource, AsyncGoogleResource
from .resources.amazon import AmazonResource, AsyncAmazonResource
from .resources.walmart import WalmartResource, AsyncWalmartResource
from .resources.youtube import YouTubeResource, AsyncYouTubeResource

DEFAULT_BASE_URL = "https://api.webdataapi.com"


class ApiClient:
    fetch: FetchResource
    jobs: JobsResource
    google: GoogleResource
    amazon: AmazonResource
    walmart: WalmartResource
    youtube: YouTubeResource

    def __init__(
        self,
        api_key: str,
        *,
        base_url: str = DEFAULT_BASE_URL,
        timeout: float = DEFAULT_TIMEOUT,
        max_retries: int = DEFAULT_MAX_RETRIES,
    ) -> None:
        http = SyncHttpClient(base_url, api_key, timeout=timeout, max_retries=max_retries)
        self.fetch = FetchResource(http)
        self.jobs = JobsResource(http)
        self.google = GoogleResource(http)
        self.amazon = AmazonResource(http)
        self.walmart = WalmartResource(http)
        self.youtube = YouTubeResource(http)
        self._http = http

    def close(self) -> None:
        self._http.close()

    def __enter__(self) -> "ApiClient":
        return self

    def __exit__(self, *args: object) -> None:
        self.close()


class AsyncApiClient:
    fetch: AsyncFetchResource
    jobs: AsyncJobsResource
    google: AsyncGoogleResource
    amazon: AsyncAmazonResource
    walmart: AsyncWalmartResource
    youtube: AsyncYouTubeResource

    def __init__(
        self,
        api_key: str,
        *,
        base_url: str = DEFAULT_BASE_URL,
        timeout: float = DEFAULT_TIMEOUT,
        max_retries: int = DEFAULT_MAX_RETRIES,
    ) -> None:
        http = AsyncHttpClient(base_url, api_key, timeout=timeout, max_retries=max_retries)
        self.fetch = AsyncFetchResource(http)
        self.jobs = AsyncJobsResource(http)
        self.google = AsyncGoogleResource(http)
        self.amazon = AsyncAmazonResource(http)
        self.walmart = AsyncWalmartResource(http)
        self.youtube = AsyncYouTubeResource(http)
        self._http = http

    async def aclose(self) -> None:
        await self._http.aclose()

    async def __aenter__(self) -> "AsyncApiClient":
        return self

    async def __aexit__(self, *args: object) -> None:
        await self.aclose()
