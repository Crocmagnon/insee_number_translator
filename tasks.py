from concurrent.futures import ThreadPoolExecutor
from pathlib import Path

from invoke import Context, task

TARGETS = [
    "aix/ppc64",
    "darwin/amd64",
    "dragonfly/amd64",
    "freebsd/386",
    "freebsd/amd64",
    "freebsd/arm",
    "freebsd/arm64",
    "illumos/amd64",
    "linux/386",
    "linux/amd64",
    "linux/arm",
    "linux/arm64",
    "linux/ppc64",
    "linux/ppc64le",
    "linux/mips",
    "linux/mipsle",
    "linux/mips64",
    "linux/mips64le",
    "linux/riscv64",
    "linux/s390x",
    "netbsd/386",
    "netbsd/amd64",
    "netbsd/arm",
    "netbsd/arm64",
    "openbsd/386",
    "openbsd/amd64",
    "openbsd/arm",
    "openbsd/arm64",
    "plan9/386",
    "plan9/amd64",
    "plan9/arm",
    "solaris/amd64",
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
def build(context):
    """Cross-platform build"""
    with ThreadPoolExecutor() as pool:
        for target in TARGETS:
            os, arch = target.split("/")
            binary_name = f"insee-{os}-{arch}"
            if os == "windows":
                binary_name += ".exe"
            binary_path = Path(__file__).resolve(strict=True).parent / "dist" / binary_name
            pool.submit(context.run, f"go build -o {binary_path}", env={"GOOS": os, "GOARCH": arch}, echo=True)
