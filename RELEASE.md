# Release

To release new version of Kuma follow the steps:

1. Create a CHANGELOG.md.
2. Double-check with changelog that all new features are documented on kuma.io website.
3. Create PR to [kuma.io website repository](https://github.com/Kong/kuma-website) with new download links and the new version.
4. Create a new git tag.
5. Push git tag. This will trigger the release job on CI.
6. Make sure that new binaries are available in [Bintray](https://bintray.com/kong/kuma).
7. Download the new Kuma version and double check that it works with demo app. Check that is works both in `universal` and `kubernetes` modes.
8. Merge PR to website repository.
9. Create a new [Github release](https://github.com/Kong/kuma/releases)
10. Announce new version on Kuma Slack #news channel.

## Major releases
For major releases make sure that you also:

1. Create a blog post on Kong's blog.
2. Send newsletter about new release.