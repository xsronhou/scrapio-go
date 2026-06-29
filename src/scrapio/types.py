from __future__ import annotations
from typing import Any, Literal, Optional, Union
from pydantic import BaseModel


# ---- Fetch ----

class FetchSession(BaseModel):
    id: str

class FetchRequest(BaseModel):
    url: str
    render_js: Optional[bool] = None
    device: Optional[Literal["desktop", "mobile", "tablet"]] = None
    session: Optional[FetchSession] = None
    output: Optional[list[str]] = None
    extract: Optional[dict[str, Any]] = None
    actions: Optional[list[Any]] = None
    timeout: Optional[int] = None
    proxy: Optional[str] = None
    country: Optional[str] = None

class FetchResponse(BaseModel):
    request_id: str
    url: str
    status_code: int
    outputs: dict[str, Any]
    diagnostics: Optional[dict[str, Any]] = None


# ---- Jobs ----

JobStatus = Literal["queued", "running", "completed", "partial", "failed", "cancelled"]

class CreateJobRequest(BaseModel):
    job_type: str
    payload: dict[str, Any]
    webhook_url: Optional[str] = None

class Job(BaseModel):
    request_id: str
    job_id: str
    job_type: str
    status: str
    created_at: str
    updated_at: Optional[str] = None
    webhook_url: Optional[str] = None

class JobError(BaseModel):
    code: str
    message: str

class JobResult(Job):
    result: Optional[Any] = None
    error: Optional[JobError] = None


# ---- Google ----

GoogleSearchType = Literal["classic", "news", "maps", "images", "lens", "shopping", "ai_mode", "ads"]
GoogleDevice = Literal["desktop", "mobile"]
GoogleDateRange = Literal["past_hour", "past_day", "past_week", "past_month", "past_year"]
GoogleSortBy = Literal["relevance", "reviews", "price_asc", "price_desc"]

class GoogleSearchParams(BaseModel):
    search: str
    search_type: Optional[GoogleSearchType] = None
    country_code: Optional[str] = None
    language: Optional[str] = None
    device: Optional[GoogleDevice] = None
    page: Optional[Union[int, str]] = None
    date_range: Optional[GoogleDateRange] = None
    latitude: Optional[Union[float, str]] = None
    longitude: Optional[Union[float, str]] = None
    radius: Optional[Union[float, str]] = None
    sort_by: Optional[GoogleSortBy] = None

class GoogleSearchResponse(BaseModel):
    request_id: str
    results: list[Any]
    pagination: Optional[dict[str, Any]] = None


# ---- Amazon ----

class AmazonProductResponse(BaseModel):
    provider: str
    asin: str
    title: str
    brand: Optional[str] = None
    price: Optional[float] = None
    currency: Optional[str] = None
    availability: Optional[str] = None
    rating: Optional[float] = None
    review_count: Optional[int] = None
    images: Optional[list[str]] = None
    bullet_points: Optional[list[str]] = None
    url: str
    model_config = {"extra": "allow"}

class AmazonSearchResponse(BaseModel):
    request_id: str
    results: list[AmazonProductResponse]


# ---- Walmart ----

class WalmartProductResponse(BaseModel):
    provider: str
    product_id: str
    title: str
    brand: Optional[str] = None
    price: Optional[float] = None
    availability: Optional[str] = None
    url: str
    model_config = {"extra": "allow"}

class WalmartSearchResponse(BaseModel):
    request_id: str
    results: list[WalmartProductResponse]


# ---- YouTube ----

class YouTubeVideoResponse(BaseModel):
    request_id: str
    video: dict[str, Any]

class YouTubeSearchResponse(BaseModel):
    request_id: str
    results: list[Any]

class YouTubeSubtitleResponse(BaseModel):
    request_id: str
    video_id: str
    subtitles: list[Any]
