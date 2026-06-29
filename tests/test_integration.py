"""
SDK integration tests — require a running API stack.

Set WDA_API_KEY and WDA_BASE_URL, then run:
    WDA_API_KEY=sk-test WDA_BASE_URL=http://localhost:3000 pytest tests/
"""

import os
import pytest
import pytest_asyncio

API_KEY = os.environ.get("WDA_API_KEY", "")
BASE_URL = os.environ.get("WDA_BASE_URL", "http://localhost:3000")

skip_if_no_key = pytest.mark.skipif(not API_KEY, reason="WDA_API_KEY not set")


# ---- Sync client tests ----

@skip_if_no_key
class TestSyncFetch:
    def setup_method(self):
        from webdataapi import ApiClient
        self.client = ApiClient(api_key=API_KEY, base_url=BASE_URL)

    def teardown_method(self):
        self.client.close()

    def test_fetch_returns_markdown(self):
        from webdataapi import FetchRequest
        res = self.client.fetch.fetch(FetchRequest(url="https://example.com", output=["markdown"]))
        assert res.request_id
        assert res.status_code == 200
        assert "markdown" in res.outputs


@skip_if_no_key
class TestSyncJobs:
    def setup_method(self):
        from webdataapi import ApiClient
        self.client = ApiClient(api_key=API_KEY, base_url=BASE_URL)

    def teardown_method(self):
        self.client.close()

    def test_create_and_poll_job(self):
        from webdataapi import CreateJobRequest
        job = self.client.jobs.create(
            CreateJobRequest(job_type="fetch", payload={"url": "https://example.com", "output": ["markdown"]})
        )
        assert job.job_id
        assert job.status in ("queued", "running", "completed", "partial")

        polled = self.client.jobs.get(job.job_id)
        assert polled.job_id == job.job_id

    def test_wait_for_completion(self):
        from webdataapi import CreateJobRequest
        job = self.client.jobs.create(
            CreateJobRequest(job_type="fetch", payload={"url": "https://example.com", "output": ["markdown"]})
        )
        result = self.client.jobs.wait_for_completion(job.job_id, poll_interval=1.0, timeout=60.0)
        assert result.status in ("completed", "partial", "failed", "cancelled")


@skip_if_no_key
class TestSyncGoogle:
    def setup_method(self):
        from webdataapi import ApiClient
        self.client = ApiClient(api_key=API_KEY, base_url=BASE_URL)

    def teardown_method(self):
        self.client.close()

    def test_search_returns_results(self):
        from webdataapi import GoogleSearchParams
        res = self.client.google.search(GoogleSearchParams(search="python web scraping"))
        assert res.request_id
        assert isinstance(res.results, list)


@skip_if_no_key
class TestSyncAmazon:
    def setup_method(self):
        from webdataapi import ApiClient
        self.client = ApiClient(api_key=API_KEY, base_url=BASE_URL)

    def teardown_method(self):
        self.client.close()

    def test_search_amazon(self):
        res = self.client.amazon.search("laptop")
        assert res.request_id
        assert isinstance(res.results, list)


@skip_if_no_key
class TestSyncWalmart:
    def setup_method(self):
        from webdataapi import ApiClient
        self.client = ApiClient(api_key=API_KEY, base_url=BASE_URL)

    def teardown_method(self):
        self.client.close()

    def test_search_walmart(self):
        res = self.client.walmart.search("headphones")
        assert res.request_id
        assert isinstance(res.results, list)


@skip_if_no_key
class TestSyncYouTube:
    def setup_method(self):
        from webdataapi import ApiClient
        self.client = ApiClient(api_key=API_KEY, base_url=BASE_URL)

    def teardown_method(self):
        self.client.close()

    def test_search_youtube(self):
        res = self.client.youtube.search("python tutorial")
        assert res.request_id
        assert isinstance(res.results, list)


# ---- Error types (sync) ----

class TestSyncErrors:
    def test_auth_error_on_bad_key(self):
        from webdataapi import ApiClient, AuthError, FetchRequest
        client = ApiClient(api_key="invalid", base_url=BASE_URL)
        with pytest.raises(AuthError) as exc_info:
            client.fetch.fetch(FetchRequest(url="https://example.com"))
        assert exc_info.value.status_code == 401

    def test_credits_exhausted_error_is_catchable(self):
        from webdataapi import CreditsExhaustedError
        body = {"request_id": "r1", "error": {"code": "credits_exhausted", "message": "no credits"}}
        err = CreditsExhaustedError(body)
        assert err.status_code == 402
        assert isinstance(err, CreditsExhaustedError)

    def test_rate_limit_error_is_catchable(self):
        from webdataapi import RateLimitError
        body = {"request_id": "r2", "error": {"code": "rate_limited", "message": "slow down"}}
        err = RateLimitError(body)
        assert err.status_code == 429
        assert isinstance(err, RateLimitError)


# ---- Async client tests ----

@skip_if_no_key
@pytest.mark.asyncio
class TestAsyncFetch:
    async def test_fetch_returns_markdown(self):
        from webdataapi import AsyncApiClient, FetchRequest
        async with AsyncApiClient(api_key=API_KEY, base_url=BASE_URL) as client:
            res = await client.fetch.fetch(FetchRequest(url="https://example.com", output=["markdown"]))
            assert res.request_id
            assert res.status_code == 200

    async def test_wait_for_completion_async(self):
        from webdataapi import AsyncApiClient, CreateJobRequest
        async with AsyncApiClient(api_key=API_KEY, base_url=BASE_URL) as client:
            job = await client.jobs.create(
                CreateJobRequest(job_type="fetch", payload={"url": "https://example.com", "output": ["markdown"]})
            )
            result = await client.jobs.wait_for_completion(job.job_id, poll_interval=1.0, timeout=60.0)
            assert result.status in ("completed", "partial", "failed", "cancelled")

    async def test_auth_error_async(self):
        from webdataapi import AsyncApiClient, AuthError, FetchRequest
        async with AsyncApiClient(api_key="invalid", base_url=BASE_URL) as client:
            with pytest.raises(AuthError):
                await client.fetch.fetch(FetchRequest(url="https://example.com"))
