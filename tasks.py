from invoke import Context, task


@task
def tag(context, tag):
    """Create & push a git tag"""
    context: Context
    context.run(f"git tag -a {tag} -m '{tag}'")
    context.run("git push --follow-tags")
