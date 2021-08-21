import re
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
BASE_DIR = Path(__file__).parent.resolve(strict=True)


@task
def test(context):
    """Run tests"""
    context: Context
    context.run(f"go test {BASE_DIR}/...", echo=True)


@task(pre=[test])
def release(context, version_name):
    """Create & push git tag + build binaries"""
    tag(context, version_name)
    build(context, version_name)


@task(pre=[test])
def tag(context, version_name):
    """Create & push a git tag"""
    context: Context
    context.run(f"git tag -a {version_name} -m '{version_name}'", echo=True)
    context.run("git push --follow-tags", echo=True)


@task
def build(context, version_name):
    """Cross-platform build"""
    with ThreadPoolExecutor() as pool:
        for target in TARGETS:
            os, arch = target.split("/")
            binary_name = f"insee-{version_name}-{os}-{arch}"
            if os == "windows":
                binary_name += ".exe"
            binary_path = BASE_DIR / "dist" / binary_name
            pool.submit(
                context.run,
                f"go build -o {binary_path}",
                env={"GOOS": os, "GOARCH": arch},
                echo=True,
            )


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
