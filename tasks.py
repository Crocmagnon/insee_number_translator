from concurrent.futures import ThreadPoolExecutor
from pathlib import Path

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


@task
def tag(context, tag):
    """Create & push a git tag"""
    context: Context
    context.run(f"git tag -a {tag} -m '{tag}'")
    context.run("git push --follow-tags")


@task
def build(context, version_name):
    """Cross-platform build"""
    with ThreadPoolExecutor() as pool:
        for target in TARGETS:
            os, arch = target.split("/")
            binary_name = f"insee-{version_name}-{os}-{arch}"
            if os == "windows":
                binary_name += ".exe"
            binary_path = (
                Path(__file__).resolve(strict=True).parent / "dist" / binary_name
            )
            pool.submit(
                context.run,
                f"go build -o {binary_path}",
                env={"GOOS": os, "GOARCH": arch},
                echo=True,
            )
