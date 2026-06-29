from __future__ import annotations
import asyncio
import time
from .._http import SyncHttpClient, AsyncHttpClient
from ..types import CreateJobRequest, Job, JobResult

TERMINAL = {"completed", "partial", "failed", "cancelled"}


class JobsResource:
    def __init__(self, http: SyncHttpClient) -> None:
        self._http = http

    def create(self, request: CreateJobRequest) -> Job:
        return self._http.request(
            "POST", "/v1/jobs",
            json=request.model_dump(exclude_none=True),
            response_model=Job,
        )

    def get(self, job_id: str) -> Job:
        return self._http.request("GET", f"/v1/jobs/{job_id}", response_model=Job)

    def get_result(self, job_id: str) -> JobResult:
        return self._http.request("GET", f"/v1/jobs/{job_id}/result", response_model=JobResult)

    def wait_for_completion(
        self,
        job_id: str,
        *,
        poll_interval: float = 2.0,
        timeout: float = 300.0,
    ) -> JobResult:
        deadline = time.monotonic() + timeout
        while time.monotonic() < deadline:
            job = self.get(job_id)
            if job.status in TERMINAL:
                return self.get_result(job_id)
            time.sleep(poll_interval)
        raise TimeoutError(f"Job {job_id} did not complete within {timeout}s")


class AsyncJobsResource:
    def __init__(self, http: AsyncHttpClient) -> None:
        self._http = http

    async def create(self, request: CreateJobRequest) -> Job:
        return await self._http.request(
            "POST", "/v1/jobs",
            json=request.model_dump(exclude_none=True),
            response_model=Job,
        )

    async def get(self, job_id: str) -> Job:
        return await self._http.request("GET", f"/v1/jobs/{job_id}", response_model=Job)

    async def get_result(self, job_id: str) -> JobResult:
        return await self._http.request("GET", f"/v1/jobs/{job_id}/result", response_model=JobResult)

    async def wait_for_completion(
        self,
        job_id: str,
        *,
        poll_interval: float = 2.0,
        timeout: float = 300.0,
    ) -> JobResult:
        deadline = asyncio.get_event_loop().time() + timeout
        while asyncio.get_event_loop().time() < deadline:
            job = await self.get(job_id)
            if job.status in TERMINAL:
                return await self.get_result(job_id)
            await asyncio.sleep(poll_interval)
        raise TimeoutError(f"Job {job_id} did not complete within {timeout}s")
