import os
import re
from concurrent.futures import ThreadPoolExecutor
from pathlib import Path
from typing import List

import requests
from invoke import Context, task

TARGETS = [
    "darwin/amd64",
    "freebsd/386",
    "freebsd/amd64",
    "freebsd/arm",
    "freebsd/arm64",
    "linux/386",
    "linux/amd64",
    "linux/arm",
    "linux/arm64",
    "windows/386",
    "windows/amd64",
    "windows/arm",
]
BASE_DIR = Path(__file__).parent.resolve(strict=True)
DIST_DIR = BASE_DIR / "dist"
GITEA_TOKEN = os.getenv("GITEA_TOKEN")


@task
def test(context):
    """Run tests"""
    context: Context
    with context.cd(BASE_DIR):
        context.run(f"go test ./... -race -bench .", echo=True)


@task
def clean(context):
    """Clean dist files"""
    context.run(f"rm -rf {DIST_DIR}", echo=True)


@task(pre=[test], post=[clean])
def release(context, version_name):
    """Create & push git tag + build binaries"""
    tag(context, version_name)
    binaries = build(context, version_name)
    upload(context, version_name, binaries)


@task(pre=[test])
def tag(context, version_name):
    """Create & push a git tag"""
    context: Context
    version_name = fix_version_name(version_name)
    context.run(f"git tag -a {version_name} -m '{version_name}'", echo=True)
    context.run("git push --follow-tags", echo=True)


@task
def build(context, version_name):
    """Cross-platform build"""
    version_name = fix_version_name(version_name)
    binaries = []
    with ThreadPoolExecutor() as pool:
        for target in TARGETS:
            os, arch = target.split("/")
            binary_name = f"insee-{version_name}-{os}-{arch}"
            if os == "windows":
                binary_name += ".exe"
            binary_path = DIST_DIR / binary_name
            binaries.append(binary_path)
            pool.submit(
                context.run,
                f"go build -o {binary_path}",
                env={"GOOS": os, "GOARCH": arch},
                echo=True,
            )
    return binaries


@task
def upload(ctx, version_name, binaries):
    context: Context
    version_name = fix_version_name(version_name)
    session = requests.Session()
    if not GITEA_TOKEN:
        raise ValueError("You need to set the GITEA_TOKEN env var before uploading")
    session.headers["Authorization"] = f"token {GITEA_TOKEN}"
    url = "https://git.augendre.info/api/v1/repos/gaugendre/insee_number_translator/releases"
    resp = session.post(
        url, json={"name": version_name, "tag_name": version_name, "draft": True}
    )
    resp.raise_for_status()
    resp = resp.json()
    html_url = resp.get("html_url")
    print(f"The draft release has been created at {html_url}")
    api_url = resp.get("url") + "/assets"
    with ThreadPoolExecutor() as pool:
        for binary in binaries:
            pool.submit(post_attachment, api_url, binary, session)
    print(f"All uploads are finished. Update & publish your draft: {html_url}")


def post_attachment(api_url, binary, session):
    binary = Path(binary)
    name = binary.name
    url = api_url + f"?name={name}"
    print(f"Uploading {name}...")
    with open(binary, "rb") as f:
        res = session.post(url, files={"attachment": f})
        status_code = res.status_code
        if status_code != 201:
            res = res.json()
            print(f"Status != 201 for {name}: {status_code} {res}")


@task
def pre_process(context):
    """Pre-process raw data into JSON"""
    files_to_rename = {
        r"commune.*\.csv": "commune.csv",
        r"departement.*\.csv": "departement.csv",
        r"pays.*\.csv": "pays.csv",
    }
    raw_data_dir = BASE_DIR / "data" / "raw_data"
    for file in raw_data_dir.iterdir():
        for reg, target_name in files_to_rename.items():
            reg = re.compile(reg)
            if reg.match(file.name):
                file.rename(raw_data_dir / target_name)

    with context.cd(BASE_DIR):
        context.run("go run ./pre_process")


def fix_version_name(version_name: str):
    if not version_name.startswith("v"):
        return f"v{version_name}"
    return version_name
