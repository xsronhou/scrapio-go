from __future__ import annotations
from typing import Optional
from .._http import SyncHttpClient, AsyncHttpClient
from ..types import YouTubeVideoResponse, YouTubeSearchResponse, YouTubeSubtitleResponse


class YouTubeResource:
    def __init__(self, http: SyncHttpClient) -> None:
        self._http = http

    def search(self, query: str, *, page: Optional[int] = None, country: Optional[str] = None, language: Optional[str] = None) -> YouTubeSearchResponse:
        return self._http.request(
            "GET", "/v1/youtube/search",
            params={"query": query, "page": page, "country": country, "language": language},
            response_model=YouTubeSearchResponse,
        )

    def get_video(self, video_id: str) -> YouTubeVideoResponse:
        return self._http.request(
            "GET", f"/v1/youtube/videos/{video_id}",
            response_model=YouTubeVideoResponse,
        )

    def get_subtitles(self, video_id: str, *, language: Optional[str] = None) -> YouTubeSubtitleResponse:
        return self._http.request(
            "GET", "/v1/youtube/subtitles",
            params={"video_id": video_id, "language": language},
            response_model=YouTubeSubtitleResponse,
        )

    def queue_search_crawl(self, query: str, *, page: Optional[int] = None) -> dict:
        from pydantic import RootModel
        from typing import Any
        class _R(RootModel[dict[str, Any]]): pass
        result = self._http.request(
            "POST", "/v1/youtube/search/crawl",
            json={"query": query, "page": page},
            response_model=_R,
        )
        return result.root


class AsyncYouTubeResource:
    def __init__(self, http: AsyncHttpClient) -> None:
        self._http = http

    async def search(self, query: str, *, page: Optional[int] = None, country: Optional[str] = None, language: Optional[str] = None) -> YouTubeSearchResponse:
        return await self._http.request(
            "GET", "/v1/youtube/search",
            params={"query": query, "page": page, "country": country, "language": language},
            response_model=YouTubeSearchResponse,
        )

    async def get_video(self, video_id: str) -> YouTubeVideoResponse:
        return await self._http.request(
            "GET", f"/v1/youtube/videos/{video_id}",
            response_model=YouTubeVideoResponse,
        )

    async def get_subtitles(self, video_id: str, *, language: Optional[str] = None) -> YouTubeSubtitleResponse:
        return await self._http.request(
            "GET", "/v1/youtube/subtitles",
            params={"video_id": video_id, "language": language},
            response_model=YouTubeSubtitleResponse,
        )

    async def queue_search_crawl(self, query: str, *, page: Optional[int] = None) -> dict:
        from pydantic import RootModel
        from typing import Any
        class _R(RootModel[dict[str, Any]]): pass
        result = await self._http.request(
            "POST", "/v1/youtube/search/crawl",
            json={"query": query, "page": page},
            response_model=_R,
        )
        return result.root
